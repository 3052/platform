package nord

import (
   "encoding/json"
   "io"
   "net/http"
   "strconv"
)

type server struct {
   Hostname string
   Technologies []struct {
      Id int
      Identifier string
   }
}

func (s *server) proxy_ssl() bool {
   for _, technology := range s.Technologies {
      if technology.Identifier == "proxy_ssl" {
         return true
      }
   }
   return false
}

type servers []server

type Byte[T any] []byte

func (s *servers) unmarshal(data Byte[servers]) error {
   return json.Unmarshal(data, s)
}

// 0 for all
func get_servers(limit int) (Byte[servers], error) {
   req, _ := http.NewRequest("", "https://api.nordvpn.com/v2/servers", nil)
   req.URL.RawQuery = "limit=" + strconv.Itoa(limit)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}
