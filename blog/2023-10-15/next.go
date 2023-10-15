package main

import (
   "154.pages.dev/protobuf"
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
)

func main() {
   var m protobuf.Message
   m.Add(1, func(m *protobuf.Message) {
      m.Add(1, func(m *protobuf.Message) {
         m.Add_Varint(16, 3)
         m.Add_String(17, "16.49.39")
         m.Add_String(19, "12")
         m.Add_Varint(64, 32)
      })
   })
   m.Add_String(2, "2ZcDwdXEVyI")
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "youtubei.googleapis.com"
   req.URL.Path = "/youtubei/v1/next"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(m.Append(nil)))
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      panic(res.Status)
   }
   data, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   if bytes.Contains(data, []byte("In The Heat Of The Night")) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}
