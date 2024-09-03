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
   address justwatch.Address
}

func main() {
   var f flags
   flag.Var(&f.address, "a", "address")
   flag.DurationVar(&f.sleep, "s", 99*time.Millisecond, "sleep")
   flag.BoolVar(&f.all, "all", false, "all results")
   flag.Parse()
   text.Transport{}.Set(true)
   if f.address.Path != "" {
      err := f.stream()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
