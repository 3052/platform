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
   req.Header["Accept-Encoding"] = []string{"identity"}
   req.Header["Cache-Control"] = []string{"no-cache, no-store, max-age=0"}
   req.Header["Client-Token"] = []string{"AABQbHRJEprCYkFDUpq8Lub8iVqhpnRoSwbSBZ7ptgT4EIuFbnbM6+wNvBu2aEXGZ/WgpqTCKt9I1r7qJBxDn341VA/tYwv7Zq0u0H0ERY46g/R6iqy/63AsGk7+Un/0+D8DIR7RspG5qQY9cHT0xy5q4DbuRo8sUU0K5h96sArhb1T2HM9/+Snadkrh10N9zhInLDKzgEUP6EwCvZYf++e4gmz5+cGVtmnRE3k1pNViPZO2h+lZ0NlUFrIIFwR7pTYxMLfKfMqi5MxDUzz57JZ0N+eGu2KDiBOwghvH+6PR68TWJnjm2IZ2HHHWs3qM1+fXwziKurWEEZtgwSAEn+J7n5kqcBg="}
   req.Header["Content-Length"] = []string{"441"}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["User-Agent"] = []string{"Spotify/8.9.68.456 Android/23 (Android SDK built for x86)"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "login5.spotify.com"
   req.URL.Path = "/v3/login"
   req.URL.RawPath = ""
   value := url.Values{}
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

var body = strings.NewReader("\n4\n 9a8d2f0ce77a4e248bb71fefcb557637\x12\x109aa96a99ce83cbb8\x12\x99\x02\x03\x00\x17N3i嫣A[\xe6KU\xfcX9\xa2\x85\xee\xc0\x8c\xcbR\n\f\xa9\x066j\xe4\x06{q\xe4\x80\x15\xe9\x82Ӻ\xf9\x05\xb1X͙\xbcj\x06vg\xf7GT\xd9`\x11\xea\xe1DA\xb6G\xc97ȓw}\xa8wε\xd9YZ\xa2\xe9\xb5n\x05a\x1a%E=\xee\x1cw\xba\xa4h\xc8s\x9d\xb1\x8bN\xc66\x122Ǔ:T<\xbd\xe0\x17Ԍ\xf2:k״L\b\xed\xf2\x91\f¹\x98G4WJ\xcc\xc1\a\x1cd\xfa\a\xcc\xf9\xabY5*\xe1\xc4\xef\xf9CO\x1a\xed\xa8ײ\xc79I\xf3<k\xe5\x89\x1bp E\x8a\xb6x\xd0|{i\xc2X\xea\xdc`\xc4\xe3o)x\x1f\xb1\x92\xa7\x9d\x98O\xc5:\xfd\xd4\xe7\u0097\x9aw\xbd\x99\f\v&\xaa#fp\xbc\x01\xfe\x91Ⳇ'\\\xf9#\xd6\xccCp!{<h\xb4\xa4\xb7ny\xaeHk\x9e\xc6r4\xe9\x87\x15(\x84I5\xcf\xe7\x90_Č-\xddq\x80\xbd\xa5\x01\xf6\x98\xa2\xd0+\x92 \n\xaaEQ`\xb8O\x80/\n2F\xfc\xb1\x1a\x1c\n\x1a\n\x18\n\x10\xf3\xf2\x97\xc7a_\x11\x93\x00\x00\x00\x00\x00\x00\f\xd6\x12\x04\x10\xa8\x9e\x1e\xaa\x06F\n\x152024-8-31@mailsac.com\x12\vBetter-Seen\x1a                                 ")
