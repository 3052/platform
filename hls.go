package stream

import (
   "154.pages.dev/http/option"
   "154.pages.dev/stream/hls"
   "fmt"
   "io"
   "net/http"
   "os"
)

func hls_get[T hls.Mixed](str Stream, items []T, index int) error {
   if str.Info {
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
   file_name, err := Name(str.Namer)
   if err != nil {
      return err
   }
   file, err := os.Create(file_name + item.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   req, err := http.NewRequest("GET", item.URI(), nil)
   if err != nil {
      return err
   }
   req.URL = str.Base.ResolveReference(req.URL)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   seg, err := hls.New_Scanner(res.Body).Segment()
   if err != nil {
      return err
   }
   var block *hls.Block
   if seg.Key != "" {
      res, err := http.Get(seg.Key)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      key, err := io.ReadAll(res.Body)
      if err != nil {
         return err
      }
      block, err = hls.New_Block(key)
      if err != nil {
         return err
      }
   }
   option.Location()
   option.Silent()
   pro := option.Progress_Parts(len(seg.URI))
   req.Host = ""
   for _, ref := range seg.URI {
      req.URL, err = str.Base.Parse(ref)
      if err != nil {
         return err
      }
      res, err := http.DefaultClient.Do(req)
      if err != nil {
         return err
      }
      if block != nil {
         data, err := io.ReadAll(pro.Reader(res))
         if err != nil {
            return err
         }
         data = block.Decrypt_Key(data)
         if _, err := file.Write(data); err != nil {
            return err
         }
      } else {
         _, err := io.Copy(file, pro.Reader(res))
         if err != nil {
            return err
         }
      }
      if err := res.Body.Close(); err != nil {
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
   s.Base = res.Request.URL
   return hls.New_Scanner(res.Body).Master()
}

