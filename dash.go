package rosso

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/sofia"
   "154.pages.dev/widevine"
   "io"
   "net/http"
   "net/url"
   "os"
   "text/template"
)

func (s Stream) key(point dash.Pointer) ([]byte, error) {
   var module widevine.CDM
   private_key, err := os.ReadFile(s.Private_Key)
   if err != nil {
      return nil, err
   }
   if err := module.New(private_key); err != nil {
      return nil, err
   }
   client_id, err := os.ReadFile(s.Client_ID)
   if err != nil {
      return nil, err
   }
   pssh, err := point.PSSH()
   if err != nil {
      kid, err := point.Default_KID()
      if err != nil {
         return nil, err
      }
      module.Key_ID(client_id, kid)
   } else {
      err := module.PSSH(client_id, pssh)
      if err != nil {
         return nil, err
      }
   }
   return module.Key(s.Poster)
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

func (s Stream) DASH_Sofia(media dash.MPD, id string) error {
   if id != "" {
      var err error
      media.Some(func(p dash.Pointer) bool {
         if p.Representation.ID == id {
            return true
         }
         ext, ok := p.Ext()
         if !ok {
            return true
         }
         initial, ok := p.Initialization()
         if ok {
            err = s.segment_template(ext, initial, p)
         } else {
            err = s.segment_base(ext, p.Representation.BaseURL, p)
         }
         return false
      })
      return err
   }
   tmpl, err := new(template.Template).Parse(dash.Template)
   if err != nil {
      return err
   }
   return tmpl.Execute(os.Stdout, media)
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
      err := sample.Decrypt_CENC(data, key)
      if err != nil {
         return err
      }
   }
   return f.Encode(dst)
}

type Stream struct {
   Base *url.URL
   Client_ID string
   Info bool
   Name string
   Poster widevine.Poster
   Private_Key string
}
