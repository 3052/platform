package main

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Accept-Encoding"] = []string{"gzip, deflate, br"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Content-Length"] = []string{"158"}
   req.Header["Content-Type"] = []string{"application/json; charset=UTF-8"}
   req.Header["Cookie"] = []string{"JSESSIONID=0280ECC990D7F7452963EC6660D39C1E", "BIGipServerpool_prod_iso_obp=914903434.36895.0000"}
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
   req.URL.Path = "/obp/ui/UIDL/"
   req.URL.RawPath = ""
   val := make(url.Values)
   val["v-uiId"] = []string{"0"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

var body = strings.NewReader(`{"csrfToken":"4f29c902-ba47-435c-ab7a-339f153c93c6","rpc":[["135","com.vaadin.shared.data.DataRequestRpc","requestRows",[0,249,0,0]]],"syncId":4,"clientId":4}`)
