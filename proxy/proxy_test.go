package proxy

import (
   "log"
   "net/http"
   "testing"
)

func Test(t *testing.T) {
   http.DefaultClient.Transport = &Transport{Proxy: http.ProxyFromEnvironment}
   log.SetFlags(log.Ltime)
   req, err := http.NewRequest("HEAD", "https://example.com", nil)
   if err != nil {
      t.Fatal(err)
   }
   _, err = http.DefaultClient.Do(req)
   if err != nil {
      t.Fatal(err)
   }
   req.URL.Path = "/proxy"
   req.Header.Set("proxy", "true")
   _, err = http.DefaultClient.Do(req)
   if err != nil {
      t.Fatal(err)
   }
}
