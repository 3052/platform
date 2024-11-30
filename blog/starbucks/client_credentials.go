package main

import (
   "crypto/md5"
   "encoding/hex"
   "io"
   "net/http"
   "net/url"
   "os"
   "strconv"
   "strings"
   "time"
)

const client_id = "QAabWxtnxV54I78dzAlviTxFhPcKsa7x"

const client_secret = "m1VDbrdVCXAZmoDK"

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.Method = "POST"
   req.URL = &url.URL{}
   req.URL.Host = "openapi.starbucks.com"
   req.URL.Path = "/v1/oauth/token"
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   req.Header["X-Api-Key"] = []string{client_id}
   value := url.Values{}
   value["sig"] = []string{func() string {
      data := []byte(client_id)
      data = append(data, client_secret...)
      data = append(data, strconv.FormatInt(time.Now().Unix(), 10)...)
      sum := md5.Sum(data)
      return hex.EncodeToString(sum[:])
   }()}
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
   "client_id":[]string{client_id},
   "client_secret":[]string{client_secret},
   "grant_type":[]string{"client_credentials"},
}.Encode())
