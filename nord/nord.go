package nord

import (
   "encoding/json"
   "net/http"
   "strconv"
)

func servers(limit int) ([]server, error) {
   req, _ := http.NewRequest("", "https://api.nordvpn.com/v1/servers", nil)
   if limit >= 1 {
      req.URL.RawQuery = "limit=" + strconv.Itoa(limit)
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var servers1 []server
   err = json.NewDecoder(resp.Body).Decode(&servers1)
   if err != nil {
      return nil, err
   }
   return servers1, nil
}

type technology struct {
   Identifier string
   Name string
}

type server struct {
   Hostname string
   Technologies []technology
}
