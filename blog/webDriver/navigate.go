package webDriver

import (
   "bytes"
   "encoding/json"
   "net/http"
   "strings"
)

// w3c.github.io/webdriver#navigate-to
func (s session_state) navigate(url string) (*http.Response, error) {
   data, err := json.Marshal(map[string]string{
      "url": url,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest("POST", address, bytes.NewReader(data))
   if err != nil {
      return nil, err
   }
   req.URL.Path += func() string {
      var b strings.Builder
      b.WriteByte('/')
      b.WriteString(s.Value.SessionId)
      b.WriteString("/url")
      return b.String()
   }()
   req.Header.Set("content-type", "application/json")
   return http.DefaultClient.Do(req)
}
