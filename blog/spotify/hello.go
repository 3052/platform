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
   req.Header["Client-Token"] = []string{"AABqCyWxaXS5fFG1qoCORUa3V+VfkiKK3JyrWUS6bMRSN5FHCK/GRHViSUKo7ww6NuSpx4Po7g2pvZvhg6qtbRc1G0nikjCZ58y13ebdUN1ZmRvDy49fQS0lsTF3ZuFdgH8IBrqOynnguEUKHk8VdPfRwYaAbMf2iZIYJ3w+HtpQh4iExMuNSwh09EZ5uDjl1p7OtiME2N8888pMDt08EXCJDUgyNiVqxf0xvSXPgIMWSFy3GYHiSOwppD0GqHPjmTUGq7OjABUGSNlaajByz0GWM62HuztcziQb367oKoLq5C/ng4kdiGEZQLubn8jpRrJA4miWgahC7IowHbGxnk1Wpvfw4aY="}
   req.Header["Content-Length"] = []string{"440"}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["User-Agent"] = []string{"Spotify/8.9.68.456 Android/23 (Android SDK built for x86)"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   
   //req.URL.Host = "login5.spotify.com"
   //req.URL.Path = "/v3/login"
   req.URL.Host = "tlshello.agwa.name"
   
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

var body = strings.NewReader("\n4\n 9a8d2f0ce77a4e248bb71fefcb557637\x12\x109aa96a99ce83cbb8\x12\x98\x02\x03\x00\a\x9c\xa1\xd9t\ued56\xabECy\x883\xab\xd85}\xfb\x8b,V\x86\xe7\x96?֭\xa5;\xb4Ji\xae\xe0\x1f\xd2j\x89~1\xb8\xb9h\"\xa3!Q#\xbd\n]\x96\x1d\x91j\x9f\x10\x86\xa7\x87\x0f\xb1J\x90\xa5\xa8\x03Bc\x19k\x8f|\xb2~O\xfb\xc2\xeb\x1cG\x06I(&\xe0\xe95\xbb!\xbe,,\xb4\xa4\x9269\xae\xd5\xc9.N\x1c\x9e\x01R\xa1\xa8\xdat\xc5$G\xf9dh\xe4\xbf\xdaՖ\xfb>\x8fJ\xa3\xc2oQZ\xab[\xa4\x9b\xd3yO\x1ei9\xae(\b]\xbb|^\x15\"\xdcl\x00\xe5\x06\x05\xb7\x92\xa7_=\x9dYMN\x99`i\xbf*\x19ɻQy\x90\xfb)\xa5\x85\x1e\xf3\x8f\x18\x82Z\x84xM\tb\x19T\xe8\x8bZ\x962\xe9\xc1gb}\xfa`\xb7\xa4e\x19\xb2\x9a\xf6\xc4-}CBWs\x99\xff\x10\t\x0f\xfc\xf7\xb3\a/\x85Eg\xeb3\xbf\xe0\xa19\xfb;\x83\xfasT\xa0\x1bC\xd5)\x01s\xf9\xa3,\xa0\xb8$dl\xc4珊!\xaf2\xc5Q\xcf\x05\xe3\xff\x04{-88\x1a\x1c\n\x1a\n\x18\n\x10\x8ag\xf5\x89\x01\xaahN\x00\x00\x00\x00\x00\x00\x00}\x12\x04\x10\xc0\x86\x02\xaa\x06F\n\x152024-8-31@mailsac.com\x12\vBetter-Seen\x1a                                 ")
