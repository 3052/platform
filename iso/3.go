package main

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func main() {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Content-Type"] = []string{"application/json; charset=UTF-8"}
   req.Header["Origin"] = []string{"https://www.iso.org"}
   req.Header["Referer"] = []string{"https://www.iso.org/obp/ui"}
   req.Header["Sec-Fetch-Dest"] = []string{"empty"}
   req.Header["Sec-Fetch-Mode"] = []string{"cors"}
   req.Header["Sec-Fetch-Site"] = []string{"same-origin"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "www.iso.org"
   req.URL.Path = "/obp/ui/UIDL/"
   val := make(url.Values)
   val["v-uiId"] = []string{"0"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Header["Cookie"] = []string{
      "BIGipServerpool_prod_iso_obp=914903434.36895.0000",
      "JSESSIONID=" + session_id,
   }
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
   `, csrf_token)
   req.Body = io.NopCloser(strings.NewReader(body))
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

const (
   csrf_token = "aa4c1468-81e0-414f-8308-3fedf7313b41"
   session_id = "62D83BAE839BB20EC53F47DC4617BD00"
)
