package rosso

import (
   "154.pages.dev/encoding/hls"
   "154.pages.dev/log"
   "errors"
   "fmt"
   "io"
   "net/http"
   "os"
)

func (s Stream) hls_get(items hls.MasterPlaylist, index int) error {
   if s.Info {
      for i, item := range items {
         fmt.Println()
         if i == index {
            fmt.Print("!")
         }
         fmt.Println(item)
      }
      return nil
   }
   item := items[index]
   file, err := os.Create(s.Name + item.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   req, err := http.NewRequest("", item.URI(), nil)
   if err != nil {
      return err
   }
   req.URL = s.Base.ResolveReference(req.URL)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   seg, err := hls.New_Scanner(res.Body).Segment()
   if err != nil {
      return err
   }
   res, err := http.Get(seg.Key)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   key, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   block, err := hls.New_Block(key)
   if err != nil {
      return err
   }
   req.Host = ""
   log.TransportDebug()
   defer log.TransportInfo()
   var src log.ProgressMeter
   src.Set(len(seg.URI))
   for _, ref := range seg.URI {
      // with HLS, the segment URL is relative to the master URL, and the
      // fragment URL is relative to the segment URL.
      req.URL, err = req.URL.Parse(ref)
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
         if block != nil {
            data, err := io.ReadAll(src.Reader(res))
            if err != nil {
               return err
            }
            data = block.Decrypt_Key(data)
            if _, err := file.Write(data); err != nil {
               return err
            }
         } else {
            _, err := file.ReadFrom(src.Reader(res))
            if err != nil {
               return err
            }
         }
         return nil
      }()
      if err != nil {
         return err
      }
   }
   return nil
}

func (s Stream) HLS_Streams(items []hls.Stream, index int) error {
   return hls_get(s, items, index)
}

func (s Stream) HLS_Media(items []hls.Media, index int) error {
   return hls_get(s, items, index)
}

func (s *Stream) HLS(ref string) (*hls.Master, error) {
   res, err := http.Get(ref)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   s.Base = res.Request.URL
   return hls.New_Scanner(res.Body).Master()
}
