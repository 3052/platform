package main

import (
   "154.pages.dev/platform/instagram"
   "flag"
   "fmt"
)

func main() {
   var address instagram.Address
   flag.Var(&address, "a", "address")
   flag.Parse()
   if address.Shortcode != "" {
      media, err := address.Media()
      if err != nil {
         panic(err)
      }
      for i, display := range media.DisplayUrls() {
         if i >= 1 {
            fmt.Println()
         }
         fmt.Println(display)
      }
   } else {
      flag.Usage()
   }
}
