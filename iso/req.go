package main

import (
   "io"
   "net/http"
   "net/url"
   "strings"
   "fmt"
   "flag"
)

func main() {
   session_id := flag.String("s", "", "session ID")
   csrf_token := flag.String("c", "", "CSRF token")
   flag.Parse()
   switch "" {
   case *csrf_token, *session_id:
      flag.Usage()
      return
   }
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "www.iso.org"
   req.URL.Path = "/obp/ui/UIDL/"
   req.URL.Scheme = "https"
   req.Header["Cookie"] = []string{
      "JSESSIONID=" + *session_id,
   }
   val := make(url.Values)
   val["v-uiId"] = []string{"0"}
   req.URL.RawQuery = val.Encode()
   body := fmt.Sprintf(`
   {
     "csrfToken": %q,
     "rpc": [
       [
         "135",
         "com.vaadin.shared.data.DataRequestRpc",
         "requestRows",
         [
           0,
           25,
           0,
           0
         ]
       ]
     ],
     "syncId": 4,
     "clientId": 4
   }
   `, *csrf_token)
   req.Body = io.NopCloser(strings.NewReader(body))
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   fmt.Println(string(text))
   if strings.Contains(string(text), "Bhutan") {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}
