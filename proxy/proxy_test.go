package proxy

import (
   "net/http"
   "os"
   "testing"
)

func proxy_test(values ...bool) error {
   http.DefaultClient.Transport = &Transport{
      Protocols: &http.Protocols{}, // github.com/golang/go/issues/25793
      Proxy: http.ProxyFromEnvironment,
   }
   for _, value := range values {
      err := get(value)
      if err != nil {
         return err
      }
   }
   return nil
}

func Test11(t *testing.T) {
   err := proxy_test(true, true)
   if err != nil {
      t.Fatal(err)
   }
}

func Test10(t *testing.T) {
   err := proxy_test(true, false)
   if err != nil {
      t.Fatal(err)
   }
}

func Test01(t *testing.T) {
   err := proxy_test(false, true)
   if err != nil {
      t.Fatal(err)
   }
}

func Test00(t *testing.T) {
   err := proxy_test(false, false)
   if err != nil {
      t.Fatal(err)
   }
}

func get(proxy bool) error {
   req, err := http.NewRequest("", "https://httpbingo.org/get", nil)
   if err != nil {
      return err
   }
   if proxy {
      req.Header.Set("proxy", "true")
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   _, err = os.Stdout.ReadFrom(resp.Body)
   if err != nil {
      return err
   }
   return nil
}
