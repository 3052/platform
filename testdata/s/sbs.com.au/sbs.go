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
   req.URL.Host = "playback.pr.sbsod.com"
   req.URL.Path = "/stream/2386967619954"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Authorization"] = []string{"Bearer eyJhbGciOiJSUzI1NiIsImprdSI6Imh0dHBzOi8vdjEuYXBpLmF1LmphbnJhaW4uY29tLzc4Nzk1NjEwLWQ0ZjYtNDdjNS04M2ZlLTdlZTNlOGM4OWRhNC9sb2dpbi9qd2siLCJraWQiOiIxNGIyMWMwNDZjNmZjYzIyYzVlYmRkNTcyMWZkZGFmZTE1YTkxOGVlIiwidHlwIjoiSldUIn0.eyJhdWQiOlsiY2ZhNzk1YmItOGIyYS00NjczLThlMzQtYjQ4MjM5MmEzZWU5Il0sImF1dGhfdGltZSI6MTczNTI1ODQ2MywiYXpwIjoiY2ZhNzk1YmItOGIyYS00NjczLThlMzQtYjQ4MjM5MmEzZWU5IiwiY2xpZW50X2lkIjoiY2ZhNzk1YmItOGIyYS00NjczLThlMzQtYjQ4MjM5MmEzZWU5IiwiZXhwIjoxNzM1MjYyMTE4LCJnbG9iYWxfc3ViIjoiY2FwdHVyZS12MTovL3Nicy5qYW5yYWluY2FwdHVyZS5jb20vZ2NmbnRmcGo2cTV1aGo5cnd2cXpoYTVkcGovdXNlci9kNmM2MmQ0NS02ZjRmLTQxMmQtYTAwZi01Y2NkMDY2OTNmMTgiLCJpYXQiOjE3MzUyNTg1MTgsImlzcyI6Imh0dHBzOi8vdjEuYXBpLmF1LmphbnJhaW4uY29tLzc4Nzk1NjEwLWQ0ZjYtNDdjNS04M2ZlLTdlZTNlOGM4OWRhNC9sb2dpbiIsImp0aSI6IjNhOWFlNjVjLWRjNGYtNGM4Yi05NzUyLTYwMzgyY2M5ZDgyZiIsInNjb3BlIjoib3BlbmlkIG9mZmxpbmVfYWNjZXNzIiwic2lkIjoiMzAzNDY2ZTMtNmEwNy00YzJlLWI1MDAtNGRiNDQ2NzM5NDc4Iiwic3ViIjoiZDZjNjJkNDUtNmY0Zi00MTJkLWEwMGYtNWNjZDA2NjkzZjE4In0.kCIiNhdvgCus06T15DSQIE49Fy9YcAPAxofu8C0k49Nywn0ymOJuTL8ojdHZmFNgwrwEIRzDgI89zxb3ijPmXtjHgTfXWzCQYlq70FwYJdRl27PogP1GmO_xO5TDtJd_FPfGfYi2Tj2IIDZza60v0P1YJlZ8tpvdZNKn--dxNlsmr-B8Sx_qGayJRKMB92X9Qb6A6AmhzNA90unlBlVRQIDUpRIKiiDaSZZIkys9YjLeRYLMi9jogBLP2BzhqjU4gOQFhk6aTvL3BbsbLCDA7crCuJQku4yy7YDy8UYmeV5OCf10_cfsn-dPjvNbRU87IwujvGn-paoWQ1qSd9ouaA"}
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

var body = strings.NewReader(`
{
   "deviceClass": "web",
   "streamProviders": [
      "HLS"
   ]
}
`)
