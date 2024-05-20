package main

import (
   "io"
   "net/http"
   "net/url"
   "strings"
   "fmt"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "gizmo.rakuten.tv"
   req.URL.Path = "/v3/avod/streamings"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/json"}
   val := make(url.Values)
   val["device_stream_video_quality"] = []string{"FHD"}
   val["device_identifier"] = []string{"atvui40"}
   req.URL.RawQuery = val.Encode()
   req.Body = io.NopCloser(body)
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   text := func() string {
      b, err := io.ReadAll(res.Body)
      if err != nil {
         panic(err)
      }
      return string(b)
   }()
   fmt.Println(text)
   if strings.Contains(text, `"video_quality":"FHD"`) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}

var body = strings.NewReader(`
{
   "player": "atvui40:DASH-CENC:WVM",
   "audio_quality": "2.0",
   "content_id": "jerry-maguire",
   "content_type": "movies",
   "device_serial": "63ac404c-08ce-4527-a9dc-3f3454d4af1b",
   "subtitle_language": "MIS",
   "video_type": "stream",
   "audio_language": "ENG",
   "classification_id": "23"
}
`)
