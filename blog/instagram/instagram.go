package main

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
   "time"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "www.instagram.com"
   req.URL.Path = "/graphql/query"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   time.Sleep(time.Second)
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

var body = strings.NewReader(url.Values{
   "doc_id":[]string{"25531498899829322"},
   "variables":[]string{`{"shortcode": "C-A8xdkCu2m"}`},
}.Encode())
