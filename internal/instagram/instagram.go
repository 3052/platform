package main

import (
   "41.neocities.org/platform/instagram"
   "41.neocities.org/text"
   "flag"
   "net/http"
   "os"
   "path"
   "time"
)

func main() {
   var code instagram.ShortCode
   flag.Var(&code, "a", "address")
   flag.Parse()
   text.Transport{}.Set(true)
   if code() != "" {
      media, err := code.Media()
      if err != nil {
         panic(err)
      }
      for _, display := range media.DisplayUrls() {
         var req http.Request
         req.URL = display()
         func() {
            resp, err := http.DefaultClient.Do(&req)
            if err != nil {
               panic(err)
            }
            defer resp.Body.Close()
            file, err := os.Create(
               path.Base(display().Path),
            )
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
