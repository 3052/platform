package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/youtube"
   "flag"
   "strconv"
   "strings"
)

type flags struct {
   itag map[int]struct{}
   r youtube.Request
   refresh bool
   request int
   v log.Level
}

func main() {
   var f flags
   f.itag = make(map[int]struct{})
   flag.Var(&f.r, "a", "address")
   flag.StringVar(&f.r.VideoId, "b", "", "video ID")
   flag.Func("i", "itag", func(s string) error {
      itag, err := strconv.Atoi(s)
      if err != nil {
         return err
      }
      f.itag[itag] = struct{}{}
      return nil
   })
   {
      var b strings.Builder
      b.WriteString("0: Android\n")
      b.WriteString("1: Android embed\n")
      b.WriteString("2: Android check")
      flag.IntVar(&f.request, "r", 0, b.String())
   }
   flag.BoolVar(&f.refresh, "refresh", false, "create OAuth refresh token")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "level")
   flag.Parse()
   log.TransportInfo()
   log.Handler(f.v)
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
