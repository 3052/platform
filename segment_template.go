package rosso

import (
   "154.pages.dev/encoding/dash"
   "154.pages.dev/log"
   "encoding/hex"
   "errors"
   "log/slog"
   "net/http"
   "os"
)

func (h HttpStream) segment_template(
   ext, initialization string, point dash.Pointer,
) error {
   key, err := h.key(point)
   if err != nil {
      return err
   }
   slog.Debug("hex", "key", hex.EncodeToString(key))
   req, err := http.NewRequest("GET", initialization, nil)
   if err != nil {
      return err
   }
   req.URL = h.Base.ResolveReference(req.URL)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   file, err := os.Create(Name(h.Name) + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   if err := encode_init(file, res.Body); err != nil {
      return err
   }
   media := point.Media()
   var meter log.ProgressMeter
   meter.Set(len(media))
   log.TransportDebug()
   defer log.TransportInfo()
   for _, ref := range media {
      // with DASH, initialization and media URLs are relative to the MPD URL
      req.URL, err = h.Base.Parse(ref)
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
         return encode_segment(file, meter.Reader(res), key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}
