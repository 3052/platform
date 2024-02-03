package rosso

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/log"
   "encoding/hex"
   "fmt"
   "log/slog"
   "net/http"
   "os"
)

func (s Stream) segment_base(
   ext, base_URL string, point dash.Pointer,
) error {
   key, err := s.key(point)
   if err != nil {
      return err
   }
   slog.Debug("*", "key", hex.EncodeToString(key))
   file, err := os.Create(s.Name + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   sb := point.Representation.SegmentBase
   // Initialization
   req, err := http.NewRequest("GET", base_URL, nil)
   if err != nil {
      return err
   }
   req.Header.Set("Range", "bytes=" + string(sb.Initialization.Range))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if err := encode_init(file, res.Body); err != nil {
      return err
   }
   byte_ranges, err := encode_sidx(base_URL, sb.IndexRange)
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
