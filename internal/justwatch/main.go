package main

import (
   "154.pages.dev/platform/justwatch"
   "154.pages.dev/text"
   "flag"
   "time"
)

type flags struct {
   filter bool
   sleep time.Duration
   v text.Level
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
   text.Transport{}.Set()
   if f.path != "" {
      err := f.stream()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
