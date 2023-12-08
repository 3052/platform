package stream

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/log"
   "154.pages.dev/stream/mp4"
   "154.pages.dev/widevine"
   "encoding/hex"
   "io"
   "log/slog"
   "net/http"
   "os"
)

func (s Stream) segment_base(ext string, item *dash.Representation) error {
   file, err := os.Create(s.Name + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   res, err := http.Get(item.BaseURL)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   index, err := item.Index()
   if err != nil {
      return err
   }
   dec := make(mp4.Decrypt)
   if err := dec.Init(io.LimitReader(res.Body, index), file); err != nil {
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
   log.Set_Transport(slog.LevelDebug)
   defer log.Set_Transport(slog.LevelInfo)
   return dec.Segment(log.New_Progress(1).Reader(res), file, key)
}
