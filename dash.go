package stream

import (
   "154.pages.dev/stream/dash"
   "154.pages.dev/stream/mp4"
   "154.pages.dev/widevine"
   "errors"
   "fmt"
   "net/http"
   "net/url"
   "os"
   option "154.pages.dev/http"
)

type Stream struct {
   Base *url.URL
   Client_ID string
   Info bool
   Name string
   Poster widevine.Poster
   Private_Key string
}

func (s Stream) DASH_Get(items []dash.Representation, index int) error {
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
   file, err := func() (*os.File, error) {
      ext, ok := item.Ext()
      if !ok {
         return nil, errors.New("extension")
      }
      return os.Create(s.Name + ext)
   }()
   if err != nil {
      return err
   }
   defer file.Close()
   req, err := func() (*http.Request, error) {
      i, ok := item.Initialization()
      if !ok {
         return nil, errors.New("initialization")
      }
      return http.NewRequest("GET", i, nil)
   }()
   if err != nil {
      return err
   }
   req.URL = s.Base.ResolveReference(req.URL)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   dec := make(mp4.Decrypt)
   if err := dec.Init(res.Body, file); err != nil {
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
   pssh, err := item.Widevine()
   if err != nil {
      return err
   }
   mod, err := widevine.New_Module(private_key, client_ID, pssh)
   if err != nil {
      return err
   }
   key, err := mod.Key(s.Poster)
   if err != nil {
      return err
   }
   option.Silent()
   media := item.Media()
   pro := option.Progress_Parts(len(media))
   for _, ref := range media {
      // with DASH, initialization and media URLs are relative to the MPD URL
      req.URL, err = s.Base.Parse(ref)
      if err != nil {
         return err
      }
      res, err := http.DefaultClient.Do(req)
      if err != nil {
         return err
      }
      if err := dec.Segment(pro.Reader(res), file, key); err != nil {
         return err
      }
      if err := res.Body.Close(); err != nil {
         return err
      }
   }
   return nil
}
