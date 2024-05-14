package main

import (
   "154.pages.dev/log"
   "154.pages.dev/platform/justwatch"
   "flag"
   "time"
)

type flags struct {
   filter bool
   sleep time.Duration
   v log.Level
   path justwatch.Path
}

func main() {
   var f flags
   flag.Var(&f.path, "a", "address")
   flag.BoolVar(&f.filter, "f", true, "filter")
   flag.DurationVar(&f.sleep, "s", 99*time.Millisecond, "sleep")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "log level")
   flag.Parse()
   f.v.Set()
   log.Transport{}.Set()
   if f.path != "" {
      err := f.stream()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
