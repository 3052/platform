package webDriver

import (
   "encoding/json"
   "net/http"
   "strings"
)

type CookieData struct {
   Value []struct {
      Name string
      Value string
   }
}

// w3c.github.io/webdriver#cookies
func (s SessionState) Cookie() (*CookieData, error) {
   req, _ := http.NewRequest("", address, nil)
   req.URL.Path += func() string {
      var b strings.Builder
      b.WriteByte('/')
      b.WriteString(s.Value.SessionId)
      b.WriteString("/cookie")
      return b.String()
   }()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   cookie := &CookieData{}
   err = json.NewDecoder(resp.Body).Decode(cookie)
   if err != nil {
      return nil, err
   }
   return cookie, nil
}
