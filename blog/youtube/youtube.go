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
   req.URL = &url.URL{}
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   req.Body = io.NopCloser(body)
   req.URL.Scheme = "https"
   req.Header["Accept"] = []string{"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"}
   req.Header["Accept-Language"] = []string{"en-us,en;q=0.5"}
   req.Header["Connection"] = []string{"close"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Host"] = []string{"www.youtube.com"}
   req.Header["Origin"] = []string{"https://www.youtube.com"}
   req.Header["Sec-Fetch-Mode"] = []string{"navigate"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3639.0 Safari/537.36"}
   req.Header["X-Youtube-Client-Name"] = []string{"7"}
   req.Header["X-Youtube-Client-Version"] = []string{"7.20241201.18.00"}
   req.Header["Cookie"] = []string{
      "GPS=1",
      "SOCS=CAI",
      "VISITOR_INFO1_LIVE=WWJS7qlmKLY",
      "YSC=io8OjZuymxw",
      "__Secure-ROLLOUT_TOKEN=CMq0r7voz5nLXBCn_v6rnoKLAxin_v6rnoKLAw%3D%3D",
      //"VISITOR_PRIVACY_METADATA=CgJVUxIEGgAgIQ%3D%3D",
   }
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

var body = strings.NewReader(`
{
 "context": {
  "client": {
   "clientName": "TVHTML5",
   "clientVersion": "7.20241201.18.00",
   "hl": "en",
   "timeZone": "UTC",
   "utcOffsetMinutes": 0
  }
 },
 "playbackContext": {
  "contentPlaybackContext": {
   "html5Preference": "HTML5_PREF_WANTS",
   "signatureTimestamp": 20102
  },
  "contentCheckOk": true,
  "racyCheckOk": true
 },
 "videoId": "Biii4l597G8"
}
`)
