package main

import (
   "154.pages.dev/platform/youtube"
   "154.pages.dev/text"
   "flag"
   "os"
   "path/filepath"
   "strings"
)

type flags struct {
   home string
   itag int
   write_code bool
   write_token bool
   read_token bool
   id youtube.VideoId
   tube youtube.InnerTube
}

func (f *flags) New() error {
   var err error
   f.home, err = os.UserHomeDir()
   if err != nil {
      return err
   }
   f.home = filepath.ToSlash(f.home)
   return nil
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.id, "a", "address")
   flag.StringVar(&f.tube.VideoId, "b", "", "video ID")
   flag.BoolVar(&f.write_code, "code", false, "write code")
   flag.IntVar(&f.itag, "i", 0, "itag")
   flag.StringVar(
      &f.tube.Context.Client.ClientName, "n",
      youtube.ClientName[0], strings.Join(youtube.ClientName[1:], "\n"),
   )
   flag.BoolVar(&f.read_token, "t", false, "read token")
   flag.BoolVar(&f.write_token, "token", false, "write token")
   flag.Parse()
   if f.tube.VideoId == "" {
      f.tube.VideoId = f.id.String()
   }
   text.Transport{}.Set(true)
   switch {
   case f.write_code:
      err := code()
      if err != nil {
         panic(err)
      }
   case f.write_token:
      err := f.token()
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
