package nord

import (
   "encoding/json"
   "net/http"
   "strconv"
)

type servers struct {
   Servers []struct {
      Hostname string
      Technologies []struct {
         Id int
      }
   }
}

func (s *servers) New(limit int) error {
   req, _ := http.NewRequest("", "https://api.nordvpn.com/v2/servers", nil)
   if limit >= 1 {
      req.URL.RawQuery = "limit=" + strconv.Itoa(limit)
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(s)
}
