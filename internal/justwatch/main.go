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
   log text.LogLevel
   path justwatch.Path
}

func main() {
   var f flags
   flag.Var(&f.path, "a", "address")
   flag.BoolVar(&f.filter, "f", true, "filter")
   flag.DurationVar(&f.sleep, "s", 99*time.Millisecond, "sleep")
   flag.TextVar(&f.log.Level, "v", f.log.Level, "log level")
   flag.Parse()
   f.log.Set()
   f.log.SetTransport(true)
   if f.path != "" {
      err := f.stream()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
