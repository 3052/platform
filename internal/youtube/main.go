package main

import (
   "154.pages.dev/platform/youtube"
   "154.pages.dev/text"
   "flag"
   "strings"
)

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.r, "a", "address")
   flag.StringVar(&f.r.VideoId, "b", "", "video ID")
   flag.BoolVar(&f.code, "c", false, "write code")
   flag.IntVar(&f.itag, "i", 0, "itag")
   {
      var b strings.Builder
      b.WriteString("0: Android\n")
      b.WriteString("1: Android embed\n")
      b.WriteString("2: Android check")
      flag.IntVar(&f.request, "r", 0, b.String())
   }
   flag.BoolVar(&f.token, "t", false, "write token")
   flag.Parse()
   text.Transport{}.Set(true)
   switch {
   case f.code:
      err := write_code()
      if err != nil {
         panic(err)
      }
   case f.token:
      err := f.write_token()
      if err != nil {
         panic(err)
      }
   case f.r.VideoId != "":
      err := f.loop()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

type flags struct {
   code bool
   home string
   itag int
   r youtube.Request
   request int
   token bool
}
