package main

import (
   "154.pages.dev/platform/youtube"
   "154.pages.dev/text"
   "flag"
   "strings"
)

type flags struct {
   itag int
   r youtube.Request
   refresh bool
   request int
}

func main() {
   var f flags
   flag.Var(&f.r, "a", "address")
   flag.StringVar(&f.r.VideoId, "b", "", "video ID")
   flag.IntVar(&f.itag, "i", 0, "itag")
   {
      var b strings.Builder
      b.WriteString("0: Android\n")
      b.WriteString("1: Android embed\n")
      b.WriteString("2: Android check")
      flag.IntVar(&f.request, "r", 0, b.String())
   }
   flag.BoolVar(&f.refresh, "refresh", false, "create OAuth refresh token")
   flag.Parse()
   text.Transport{}.Set(true)
   switch {
   case f.r.VideoId != "":
      err := f.loop()
      if err != nil {
         panic(err)
      }
   case f.refresh:
      err := f.do_refresh()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
