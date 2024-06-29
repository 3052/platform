package main

import (
   "154.pages.dev/platform/justwatch"
   "154.pages.dev/text"
   "flag"
   "time"
)

type flags struct {
   all bool
   sleep time.Duration
   path justwatch.Path
}

func main() {
   var f flags
   flag.Var(&f.path, "a", "address")
   flag.DurationVar(&f.sleep, "s", 99*time.Millisecond, "sleep")
   flag.BoolVar(&f.all, "all", false, "all results")
   flag.Parse()
   text.Transport{}.Set(true)
   if f.path != "" {
      err := f.stream()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
