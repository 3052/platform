package main

import (
   "154.pages.dev/protobuf"
   "blog/youtube"
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
)

func main() {
   var req http.Request
   req.URL = &url.URL{}
   req.URL.Path = "/youtubei/v1/get_watch"
   req.Header = http.Header{}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL.Host = "youtubei.googleapis.com"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(youtube.Request.Marshal()))
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["Priority"] = []string{"u=0, i"}
   req.Header["User-Agent"] = []string{"com.google.android.youtube/19.33.35(Linux; U; Android 9; en_US; Android SDK built for x86 Build/PSR1.180720.012) gzip"}
   req.Header["X-Goog-Api-Format-Version"] = []string{"2"}
   req.Header["X-Goog-Visitor-Id"] = []string{"Cgt0OV94ZHc0Q1hnQSjUja-2BjIKCgJVUxIEGgAgTzoMCAEgwKfTncLa8eVm"}
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      panic(err)
   }
   message := protobuf.Message{}
   err = message.Unmarshal(data)
   if err != nil {
      panic(err)
   }
   message, _ = message.Get(1)()
   message, _ = message.Get(2)()
   message, _ = message.Get(4)()
   message, _ = message.Get(3)()
   address, _ := message.GetBytes(2)()
   fmt.Println(string(address))
}
