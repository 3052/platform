package mullvad

import (
   "encoding/json"
   "errors"
   "iter"
   "log"
   "net/http"
   "os/exec"
   "strings"
   "time"
)

func command(name string, arg ...string) *exec.Cmd {
   cmd := exec.Command(name, arg...)
   log.Print(cmd.Args)
   return cmd
}

func (p *PublicRelays) OpenVpn() error {
   resp, err := http.Get("https://api.mullvad.net/public/relays/v1")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(p)
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

var ErrTimeout = errors.New("timeout")

func status(prefix string) error {
   after := time.After(30 * time.Second)
   for {
      select {
      case <-time.After(time.Second):
         data, err := command("mullvad", "status").Output()
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

func Disconnect() error {
   err := command("mullvad", "disconnect").Run()
   if err != nil {
      return err
   }
   return status("Disconnected")
}

func Connect() error {
   err := command("mullvad", "connect").Run()
   if err != nil {
      return err
   }
   return status("Connected")
}

func Relay(location ...string) error {
   arg := []string{"relay", "set", "location"}
   arg = append(arg, location...)
   data, err := command("mullvad", arg...).CombinedOutput()
   if err != nil {
      return errors.New(string(data))
   }
   return nil
}

type Transport struct{}

func (Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   err := Connect()
   if err != nil {
      return nil, err
   }
   defer Disconnect()
   log.Println(req.Method, req.URL)
   return http.DefaultTransport.RoundTrip(req)
}
