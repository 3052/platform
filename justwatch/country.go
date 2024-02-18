package main

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

const fetcher_query = `
query BackendConstantsFetcherQuery($language: Language!) {
   locales {
      fullLocale
      countryName(language: $language)
   }
}
`

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "apis.justwatch.com"
   req.URL.Path = "/graphql"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Device-Id"] = []string{
      base64.RawStdEncoding.EncodeToString(make([]byte, 16)),
   }
   body := fmt.Sprintf(`
   {
      "query": %q,
      "variables": {
         "language": "en-US"
      }
   }
   `, fetcher_query)
   req.Body = io.NopCloser(strings.NewReader(body))
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   src, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   var dst bytes.Buffer
   json.Indent(&dst, src, "", " ")
   dst.WriteTo(os.Stdout)
   if strings.Contains(string(src), `"Argentina"`) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}
