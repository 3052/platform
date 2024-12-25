package main

import (
   "flag"
   "fmt"
   "net/http"
)

type transport struct{}

func (transport) RoundTrip(req *http.Request) (*http.Response, error) {
   fmt.Println(req.URL)
   return http.DefaultTransport.RoundTrip(req)
}

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
