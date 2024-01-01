package stream

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/log"
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "log/slog"
   "net/http"
   "os"
)

func (s Stream) segment_base(
   ext, base_URL string, item *dash.Representation,
) error {
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
   file, err := os.Create(s.Name + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   sidx, moof, err := item.Sidx_Moof()
   if err != nil {
      return err
   }
   req, err := http.NewRequest("GET", base_URL, nil)
   if err != nil {
      return err
   }
   req.Header.Set("Range", "bytes=0-" + fmt.Sprint(sidx-1))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if err := encode_init(file, res.Body); err != nil {
      return err
   }
   byte_ranges, err := encode_sidx(file, base_URL, sidx, moof)
   if err != nil {
      return err
   }
   src := log.New_Progress(len(byte_ranges))
   log.Set_Transport(slog.LevelDebug)
   defer log.Set_Transport(slog.LevelInfo)
   for _, r := range byte_ranges {
      err := func() error {
         req, err := http.NewRequest("GET", base_URL, nil)
         if err != nil {
            return err
         }
         req.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", r[0], r[1]))
         res, err := http.DefaultClient.Do(req)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         return encode_segment(file, src.Reader(res), key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}
