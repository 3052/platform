package main

import (
   "154.pages.dev/platform/instagram"
   "154.pages.dev/text"
   "flag"
   "net/http"
   "os"
   "path"
   "time"
)

func main() {
   var address instagram.Address
   flag.Var(&address, "a", "address")
   flag.Parse()
   text.Transport{}.Set(true)
   if address.Shortcode != "" {
      media, err := address.Media()
      if err != nil {
         panic(err)
      }
      for _, display := range media.DisplayUrls() {
         var req http.Request
         req.URL = &display.Url
         func() {
            resp, err := http.DefaultClient.Do(&req)
            if err != nil {
               panic(err)
            }
            defer resp.Body.Close()
            file, err := os.Create(path.Base(display.Url.Path))
            if err != nil {
               panic(err)
            }
            defer file.Close()
            _, err = file.ReadFrom(resp.Body)
            if err != nil {
               panic(err)
            }
         }()
         time.Sleep(99 * time.Millisecond)
      }
   } else {
      flag.Usage()
   }
}