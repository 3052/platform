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
   req.Header = http.Header{}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["User-Agent"] = []string{"com.google.android.youtube/18.18.38(Linux; U; Android 12; en_US; sdk_gphone64_x86_64 Build/SE1B.240122.005) gzip"}
   req.Header["X-Goog-Api-Format-Version"] = []string{"2"}
   req.Header["X-Goog-Visitor-Id"] = []string{"CgtzZjI3dkR5Z0VSVSjGla22BjIKCgJVUxIEGgAgYzoMCAEgtO6MnOXY0uVm"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "youtubei.googleapis.com"
   req.URL.Path = "/youtubei/v1/player"
   value := url.Values{}
   value["id"] = []string{"40wkJJXfwQ0"}
   value["key"] = []string{"AIzaSyA8eiZmM1FaDVjRy-df2KTyQ_vz_yYM39w"}
   value["t"] = []string{"pY_UR-mBysO5"}
   req.URL.RawQuery = value.Encode()
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(youtube.Request.Marshal()))
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
   message, _ = message.Get(4)()
   messages := message.Get(3)
   for {
      message, ok := messages()
      if !ok {
         break
      }
      if v, ok := message.GetVarint(1)(); ok {
         fmt.Println(v)
      }
      if v, ok := message.GetBytes(2)(); ok {
         fmt.Println(string(v))
      }
      fmt.Println()
   }
}
