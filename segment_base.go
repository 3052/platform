package stream

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/log"
   "154.pages.dev/sofia"
   "154.pages.dev/widevine"
   "encoding/hex"
   "errors"
   "io"
   "log/slog"
   "net/http"
   "os"
)

func (s Stream) segment_template_sofia(
   ext, initialization string, item *dash.Representation,
) error {
   file, err := os.Create(s.Name + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   req, err := http.NewRequest("GET", initialization, nil)
   if err != nil {
      return err
   }
   req.URL = s.Base.ResolveReference(req.URL)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if err := encode_init(file, res.Body); err != nil {
      return err
   }
   private_key, err := os.ReadFile(s.Private_Key)
   if err != nil {
      return err
   }
   client_ID, err := os.ReadFile(s.Client_ID)
   if err != nil {
      return err
   }
   kid, err := item.Default_KID()
   if err != nil {
      return err
   }
   pssh, err := item.PSSH()
   if err != nil {
      return err
   }
   mod, err := widevine.New_Module(private_key, client_ID, kid, pssh)
   if err != nil {
      return err
   }
   key, err := mod.Key(s.Poster)
   if err != nil {
      return err
   }
   slog.Debug("*", "key", hex.EncodeToString(key))
   media, ok := item.Media()
   if !ok {
      return errors.New("Media")
   }
   src := log.New_Progress(len(media))
   log.Set_Transport(slog.LevelDebug)
   defer log.Set_Transport(slog.LevelInfo)
   for _, ref := range media {
      // with DASH, initialization and media URLs are relative to the MPD URL
      req.URL, err = s.Base.Parse(ref)
      if err != nil {
         return err
      }
      err := func() error {
         res, err := http.DefaultClient.Do(req)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         if res.StatusCode != http.StatusOK {
            return errors.New(res.Status)
         }
         return encode_segment(file, src.Reader(res), key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}

func encode_segment(dst io.Writer, src io.Reader, key []byte) error {
   var f sofia.File
   if err := f.Decode(src); err != nil {
      return err
   }
   for i := range f.Mdat.Data {
      sample := f.Mdat.Data[i]
      enc := f.Moof.Traf.Senc.Samples[i]
      err := sofia.CryptSampleCenc(sample, key, enc)
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
      if b.Header.Type() == "pssh" {
         copy(b.Header.RawType[:], "free") // Firefox
      }
   }
   stsd := &f.Moov.Trak.Mdia.Minf.Stbl.Stsd
   copy(stsd.Enca.Header.RawType[:], "mp4a") // Firefox
   copy(stsd.Encv.Header.RawType[:], "avc1") // Firefox
   for _, b := range stsd.Enca.Boxes {
      if b.Header.Type() == "sinf" {
         copy(b.Header.RawType[:], "free") // Firefox
      }
   }
   for _, b := range stsd.Encv.Boxes {
      if b.Header.Type() == "sinf" {
         copy(b.Header.RawType[:], "free") // Firefox
      }
   }
   return f.Encode(dst)
}
