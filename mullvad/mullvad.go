package mullvad

import (
   "encoding/json"
   "iter"
   "net/http"
   "os/exec"
   "strings"
   "time"
)

func status(prefix string) error {
   for range time.NewTicker(time.Second).C {
      data, err := exec.Command("mullvad", "status").Output()
      if err != nil {
         return err
      }
      text := string(data)
      if strings.HasPrefix(text, prefix) {
         if strings.Contains(text, " IPv4:") {
            break
         }
      }
   }
   return nil
}

type ServerList struct {
   Countries []struct {
      Name string
      Cities []struct {
         Name string
         Relays []struct {
            Hostname string
         }
      }
   }
}

func Disconnect() error {
   err := exec.Command("mullvad", "disconnect").Run()
   if err != nil {
      return err
   }
   return status("Disconnected")
}

func (s ServerList) Seq(country string) iter.Seq[string] {
   return func(yield func(string) bool) {
      for _, country_struct := range s.Countries {
         if country_struct.Name != country {
            continue
         }
         for _, city := range country_struct.Cities {
            for _, relay := range city.Relays {
               if !yield(relay.Hostname) {
                  return
               }
            }
         }
      }
   }
}

func Connect(location string) error {
   err := exec.Command("mullvad", "relay", "set", "location", location).Run()
   if err != nil {
      return err
   }
   err = exec.Command("mullvad", "connect").Run()
   if err != nil {
      return err
   }
   return status("Connected")
}

func (s *ServerList) OpenVpn() error {
   resp, err := http.Get("https://api.mullvad.net/public/relays/v1")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(s)
}

func (s *ServerList) WireGuard() error {
   resp, err := http.Get("https://api.mullvad.net/public/relays/wireguard/v1")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(s)
}
