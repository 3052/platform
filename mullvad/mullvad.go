package mullvad

import (
   "encoding/json"
   "errors"
   "iter"
   "net/http"
   "os/exec"
   "strings"
   "time"
)

var ErrTimeout = errors.New("timeout")

func status(prefix string) error {
   after := time.After(30 * time.Second)
   for {
      select {
      case <-time.After(time.Second):
         data, err := exec.Command("mullvad", "status").Output()
         if err != nil {
            return err
         }
         data1 := string(data)
         if strings.HasPrefix(data1, prefix) {
            if strings.Contains(data1, " IPv4:") {
               return nil
            }
         }
      case <-after:
         return ErrTimeout
      }
   }
}

func Connect(location string) error {
   data, err := exec.Command(
      "mullvad", "relay", "set", "location", location,
   ).CombinedOutput()
   if err != nil {
      return errors.New(string(data))
   }
   err = exec.Command("mullvad", "connect").Run()
   if err != nil {
      return err
   }
   return status("Connected")
}

func (p *PublicRelays) OpenVpn() error {
   resp, err := http.Get("https://api.mullvad.net/public/relays/v1")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(p)
}

func Disconnect() error {
   err := exec.Command("mullvad", "disconnect").Run()
   if err != nil {
      return err
   }
   return status("Disconnected")
}

type PublicRelays struct {
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

func (p PublicRelays) Hostname(country string) iter.Seq[string] {
   return func(yield func(string) bool) {
      for _, country1 := range p.Countries {
         if country1.Name == country {
            for _, city := range country1.Cities {
               for _, relay := range city.Relays {
                  if !yield(relay.Hostname) {
                     return
                  }
               }
            }
         }
      }
   }
}

func (p *PublicRelays) WireGuard() error {
   resp, err := http.Get("https://api.mullvad.net/public/relays/wireguard/v1")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(p)
}
