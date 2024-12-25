package main

import "flag"

type flags struct {
   address string
   location string
}

func main() {
   var f flags
   flag.StringVar(&f.address, "a", "", "address")
   flag.StringVar(&f.location, "n", "", "location")
   flag.Parse()
   if f.address != "" {
      err := f.relay()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
