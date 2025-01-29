package webDriver

import (
   "bytes"
   "encoding/json"
   "net/http"
)

const address = "http://127.0.0.1:4444/session"

// w3c.github.io/webdriver#sessions
type SessionState struct {
   Value struct {
      SessionId string
   }
}

func (s *SessionState) New() error {
   data, err := json.Marshal(map[string]any{
      "capabilities": map[string]any{},
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
