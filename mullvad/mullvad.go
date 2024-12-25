package mullvad

import (
   "encoding/json"
   "iter"
   "net/http"
   "os/exec"
   "strings"
   "time"
)

func (p *public_relays) New() error {
   resp, err := http.Get("https://api.mullvad.net/public/relays/v1")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(p)
}

func (p public_relays) relays(country string) iter.Seq[string] {
   return func(yield func(string) bool) {
      for _, country_struct := range p.Countries {
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

type public_relays struct {
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

type am_i_mullvad struct {
   MullvadExitIp bool `json:"mullvad_exit_ip"`
}

func (a *am_i_mullvad) New() error {
   resp, err := http.Get("https://ipv4.am.i.mullvad.net/json")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(a)
}

///

func status(prefix string) error {
   for range time.NewTicker(99 * time.Millisecond).C {
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

func connect(location string) error {
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

func disconnect() error {
   err := exec.Command("mullvad", "disconnect").Run()
   if err != nil {
      return err
   }
   return status("Disconnected")
}
