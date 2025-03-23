package proxy

import (
   "log"
   "net/http"
   "net/url"
)

var proxy func(*http.Request) (*url.URL, error)

type Transport http.Transport

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   if req.Header.Get("proxy") != "" {
      if t.Proxy != nil {
         if proxy == nil {
            log.Print("proxy connect")
            proxy = t.Proxy
         }
      } else if proxy != nil {
         log.Print("proxy connect")
         t.Proxy = proxy
      }
   } else if t.Proxy != nil {
      log.Print("proxy disconnect")
      proxy = t.Proxy
      t.Proxy = nil
   }
   if req.Header.Get("silent") == "" {
      log.Println(req.Method, req.URL)
   }
   return (*http.Transport)(t).RoundTrip(req)
}
