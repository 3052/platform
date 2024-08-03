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
   flag.BoolVar(&f.code, "c", false, "write code")
   flag.IntVar(&f.itag, "i", 0, "itag")
   flag.BoolVar(&f.token, "t", false, "write token")
   flag.StringVar(
      &f.tube.Context.Client.ClientName, "n",
      youtube.ClientName[0], strings.Join(youtube.ClientName[1:], " "),
   )
   flag.StringVar(&f.tube.VideoId, "b", "", "video ID")
   flag.Var(&f.id, "a", "address")
   flag.Parse()
   if f.tube.VideoId == "" {
      f.tube.VideoId = f.id.String()
   }
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
   case f.tube.VideoId != "":
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
   request int
   token bool
   tube youtube.InnerTube
   id youtube.VideoId
}
