package stream

import (
   "154.pages.dev/dash"
   "154.pages.dev/widevine"
   "errors"
   "fmt"
   "net/url"
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
   ext, ok := item.Ext()
   if !ok {
      return errors.New("extension")
   }
   initialization, ok := item.Initialization()
   if ok {
      return s.segment_template(ext, initialization, item)
   }
   return s.segment_base(ext, item)
}
