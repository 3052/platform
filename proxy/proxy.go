package proxy

import (
   "log"
   "net/http"
)

type Transport http.Transport

// need value receiver
func (t Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   if req.Header.Get("proxy") != "" {
      log.Print("proxy connect")
   } else {
      log.Print("proxy disconnect")
      t.Proxy = nil
   }
   if req.Header.Get("silent") == "" {
      log.Println(req.Method, req.URL)
   }
   return (*http.Transport)(&t).RoundTrip(req)
}
