package stream

import (
   "154.pages.dev/dash"
   "154.pages.dev/stream/mp4"
   "154.pages.dev/widevine"
   "errors"
   "fmt"
   "net/http"
   "os"
   option "154.pages.dev/http"
)

func (s Stream) segment_template(
   ext, initialization string, item dash.Representation,
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
   media, ok := item.Media()
   if !ok {
      return errors.New("Media")
   }
   pro := option.Progress_Parts(len(media))
   f := option.Silent()
   defer f()
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
            return fmt.Errorf("%v %v", res.Status, req.URL)
         }
         return dec.Segment(pro.Reader(res), file, key)
      }()
      if err != nil {
         return err
      }
   }
   return nil
}
