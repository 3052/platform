package stream

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/log"
   "154.pages.dev/widevine"
   "encoding/hex"
   "errors"
   "log/slog"
   "net/http"
   "os"
)

func (s Stream) segment_template(
   ext, initialization string, item *dash.Representation,
) error {
   file, err := os.Create(s.Name + ext)
   if err != nil {
      return err
   }
   defer file.Close()
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
