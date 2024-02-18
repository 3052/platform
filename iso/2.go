package main

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func main() {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   req.Header["Origin"] = []string{"https://www.iso.org"}
   req.Header["Referer"] = []string{"https://www.iso.org/obp/ui"}
   req.Header["Sec-Fetch-Dest"] = []string{"empty"}
   req.Header["Sec-Fetch-Mode"] = []string{"cors"}
   req.Header["Sec-Fetch-Site"] = []string{"same-origin"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "www.iso.org"
   req.URL.Path = "/obp/ui"
   val := make(url.Values)
   val["v-1708217027070"] = []string{""}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Cookie"] = []string{
      "BIGipServerpool_prod_iso_obp=914903434.36895.0000",
      "JSESSIONID=" + session_id,
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

var body = strings.NewReader(url.Values{
"theme":[]string{"iso-red"}, "v-appId":[]string{"obpui-105541713"},
"v-browserDetails":[]string{"1"}, "v-ch":[]string{"622"},
"v-curdate":[]string{"1708217027070"}, "v-cw":[]string{"1192"},
"v-dstd":[]string{"60"}, "v-dston":[]string{"false"},
"v-loc":[]string{"https://www.iso.org/obp/ui#search/code"},
"v-rtzo":[]string{"360"}, "v-sh":[]string{"864"}, "v-sw":[]string{"1536"},
"v-tzid":[]string{"America/Chicago"}, "v-tzo":[]string{"360"},
"v-vh":[]string{"0"}, "v-vw":[]string{"1192"},
"v-wn":[]string{"obpui-105541713-0.8254629619239288"},
}.Encode())
