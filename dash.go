package rosso

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/sofia"
   "154.pages.dev/widevine"
   "encoding/xml"
   "errors"
   "io"
   "net/http"
   "net/url"
   "os"
   "text/template"
)

func (h HttpStream) DASH(media *dash.MPD, id string) error {
   var point dash.Pointer
   ok := media.Contains(func(p dash.Pointer) bool {
      if p.Representation.ID == id {
         point = p
         return true
      }
      return false
   })
   if ok {
      ext, ok := point.Ext()
      if !ok {
         return errors.New("dash.Pointer.Ext")
      }
      if initial, ok := point.Initialization(); ok {
         return h.segment_template(ext, initial, point)
      }
      return h.segment_base(ext, point.Representation.BaseURL, point)
   }
   line, err := new(template.Template).Parse(dash.ModeLine)
   if err != nil {
      return err
   }
   return line.Execute(os.Stdout, media)
}

// wikipedia.org/wiki/Dynamic_Adaptive_Streaming_over_HTTP
// wikipedia.org/wiki/HTTP_Live_Streaming
type HttpStream struct {
   Base *url.URL
   Client_ID string
   Name Namer
   Poster widevine.Poster
   Private_Key string
}

func (h *HttpStream) DashMedia(uri string) (*dash.MPD, error) {
   res, err := http.Get(uri)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   h.Base = res.Request.URL
   media := new(dash.MPD)
   if err := xml.NewDecoder(res.Body).Decode(media); err != nil {
      return nil, err
   }
   return media, nil
}

func (h HttpStream) key(point dash.Pointer) ([]byte, error) {
   var pssh widevine.PSSH
   data, err := point.PSSH()
   if err != nil {
      key_id, err := point.Default_KID()
      if err != nil {
         return nil, err
      }
      pssh.Key_ID = key_id
   } else {
      err := pssh.New(data)
      if err != nil {
         return nil, err
      }
   }
   private_key, err := os.ReadFile(h.Private_Key)
   if err != nil {
      return nil, err
   }
   client_id, err := os.ReadFile(h.Client_ID)
   if err != nil {
      return nil, err
   }
   module, err := pssh.CDM(private_key, client_id)
   if err != nil {
      return nil, err
   }
   license, err := module.License(h.Poster)
   if err != nil {
      return nil, err
   }
   key, ok := module.Key(license)
   if !ok {
      return nil, errors.New("widevine.CDM.Key")
   }
   return key, nil
}

func encode_sidx(base_URL string, index dash.Range) ([][2]uint32, error) {
   req, err := http.NewRequest("GET", base_URL, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Range", "bytes=" + string(index))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var f sofia.File
   if err := f.Decode(res.Body); err != nil {
      return nil, err
   }
   _, sidx, err := index.Scan()
   if err != nil {
      return nil, err
   }
   return f.SegmentIndex.ByteRanges(uint32(sidx)+1), nil
}

func encode_init(dst io.Writer, src io.Reader) error {
   var f sofia.File
   if err := f.Decode(src); err != nil {
      return err
   }
   for _, b := range f.Movie.Boxes {
      if b.BoxHeader.BoxType() == "pssh" {
         copy(b.BoxHeader.Type[:], "free") // Firefox
      }
   }
   sd := &f.Movie.Track.Media.MediaInformation.SampleTable.SampleDescription
   if as := sd.AudioSample; as != nil {
      copy(as.ProtectionScheme.BoxHeader.Type[:], "free") // Firefox
      copy(
         as.Entry.BoxHeader.Type[:],
         as.ProtectionScheme.OriginalFormat.DataFormat[:],
      ) // Firefox
   }
   if vs := sd.VisualSample; vs != nil {
      copy(vs.ProtectionScheme.BoxHeader.Type[:], "free") // Firefox
      copy(
         vs.Entry.BoxHeader.Type[:],
         vs.ProtectionScheme.OriginalFormat.DataFormat[:],
      ) // Firefox
   }
   return f.Encode(dst)
}

func encode_segment(dst io.Writer, src io.Reader, key []byte) error {
   var f sofia.File
   if err := f.Decode(src); err != nil {
      return err
   }
   for i, data := range f.MediaData.Data {
      sample := f.MovieFragment.TrackFragment.SampleEncryption.Samples[i]
      err := sample.DecryptCenc(data, key)
      if err != nil {
         return err
      }
   }
   return f.Encode(dst)
}
