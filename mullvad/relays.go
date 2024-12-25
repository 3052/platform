package mullvad

import (
   "encoding/json"
   "iter"
   "net/http"
)

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

func (p *public_relays) New() error {
   resp, err := http.Get("https://api.mullvad.net/public/relays/v1")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(p)
}
