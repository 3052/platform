package stream

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/sofia"
   "154.pages.dev/widevine"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/url"
)

func encode_segment(dst io.Writer, src io.Reader, key []byte) error {
   var f sofia.File
   if err := f.Decode(src); err != nil {
      return err
   }
   for i, sample := range f.Mdat.Data {
      err := f.Moof.Traf.Senc.Samples[i].Decrypt_CENC(sample, key)
      if err != nil {
         return err
      }
   }
   return f.Encode(dst)
}

func encode_init(dst io.Writer, src io.Reader) error {
   var f sofia.File
   if err := f.Decode(src); err != nil {
      return err
   }
   for _, b := range f.Moov.Boxes {
      if b.Header.BoxType() == "pssh" {
         copy(b.Header.Type[:], "free") // Firefox
      }
   }
   stsd := &f.Moov.Trak.Mdia.Minf.Stbl.Stsd
   copy(stsd.Enca.Header.Type[:], "mp4a") // Firefox
   copy(stsd.Encv.Header.Type[:], "avc1") // Firefox
   for _, b := range stsd.Enca.Boxes {
      if b.Header.BoxType() == "sinf" {
         copy(b.Header.Type[:], "free") // Firefox
      }
   }
   for _, b := range stsd.Encv.Boxes {
      if b.Header.BoxType() == "sinf" {
         copy(b.Header.Type[:], "free") // Firefox
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
func (s Stream) DASH_Sofia(items []*dash.Representation, index int) error {
   if s.Info {
      for i, item := range items {
         fmt.Println()
         if i == index {
            fmt.Print("!")
         }
         fmt.Println(item)
      }
   } else if index >= 0 {
      item := items[index]
      ext, ok := item.Ext()
      if !ok {
         return errors.New("extension")
      }
      initialization, ok := item.Initialization()
      if ok {
         return s.segment_template(ext, initialization, item)
      }
      return s.segment_base(ext, item.BaseURL, item)
   }
   return nil
}

func decode_sidx(base_URL string, sidx, moof uint32) ([][2]uint32, error) {
   req, err := http.NewRequest("GET", base_URL, nil)
   if err != nil {
      return nil, err
   }
   req.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", sidx, moof-1))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var f sofia.File
   if err := f.Decode(res.Body); err != nil {
      return nil, err
   }
   return f.Sidx.Byte_Ranges(moof), nil
}
