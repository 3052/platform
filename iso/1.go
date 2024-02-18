package main

import (
   "net/http"
   "net/url"
   "os"
)

func main() {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Content-Length"] = []string{"0"}
   req.Header["Sec-Fetch-Dest"] = []string{"document"}
   req.Header["Sec-Fetch-Mode"] = []string{"navigate"}
   req.Header["Sec-Fetch-Site"] = []string{"none"}
   req.Header["Sec-Fetch-User"] = []string{"?1"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["Upgrade-Insecure-Requests"] = []string{"1"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"}
   req.Method = "HEAD"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "www.iso.org"
   req.URL.Path = "/obp/ui"
   req.URL.Scheme = "https"
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
