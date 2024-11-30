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
   req.Header = http.Header{}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "openapi.starbucks.com"
   req.URL.Path = "/v1/oauth/token"
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded; charset=UTF-8"}
   req.Header["X-Api-Key"] = []string{"QAabWxtnxV54I78dzAlviTxFhPcKsa7x"}
   value := url.Values{}
   value["sig"] = []string{"9cb4c86fd937de259ef6958faab5cbab"}
   req.URL.RawQuery = value.Encode()
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

var body = strings.NewReader(url.Values{
   "client_id":[]string{"QAabWxtnxV54I78dzAlviTxFhPcKsa7x"},
   "client_secret":[]string{"m1VDbrdVCXAZmoDK"},
   "grant_type":[]string{"client_credentials"},
}.Encode())
