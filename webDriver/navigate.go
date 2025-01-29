package webDriver

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "strings"
)

// w3c.github.io/webdriver#navigate-to
func (s SessionState) Navigate(url string) error {
   data, err := json.Marshal(map[string]string{
      "url": url,
   })
   if err != nil {
      return err
   }
   req, err := http.NewRequest("POST", address, bytes.NewReader(data))
   if err != nil {
      return err
   }
   req.URL.Path += func() string {
      var b strings.Builder
      b.WriteByte('/')
      b.WriteString(s.Value.SessionId)
      b.WriteString("/url")
      return b.String()
   }()
   req.Header.Set("content-type", "application/json")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   return resp.Write(io.Discard)
}
