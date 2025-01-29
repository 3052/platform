package webDriver

import (
   "bytes"
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
const address = "http://127.0.0.1:4444/session"

// w3c.github.io/webdriver#sessions
type SessionState struct {
   Value struct {
      SessionId string
   }
}

func (s *SessionState) New() error {
   data, err := json.Marshal(map[string]any{
      "capabilities": map[string]any{
         "alwaysMatch": map[string]any{
            "proxy": map[string]string{
               "proxyType": "manual",
               "sslProxy": "res.proxy-seller.com:10000",
            },
         },
      },
   })
   if err != nil {
      return err
   }
   req, err := http.NewRequest("POST", address, bytes.NewReader(data))
   if err != nil {
      return err
   }
   req.Header.Set("content-type", "application/json")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(s)
}
