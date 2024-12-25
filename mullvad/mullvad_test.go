package mullvad

import (
   "encoding/json"
   "fmt"
   "net/http"
   "testing"
)

const paramount_plus = "https://www.paramountplus.com/apps-api/v2.0/androidphone/video/cid/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ.json?at=ABAAAAAAAAAAAAAAAAAAAAAAU%2FEmq2DGAxNGYi71fZdi2mb6UTU2%2BelcD3bGEhs6bo4%3D"

func TestParamount(t *testing.T) {
   var relays public_relays
   err := relays.New()
   if err != nil {
      t.Fatal(err)
   }
   defer disconnect()
   for relay := range relays.relays("Canada") {
      err = connect(relay)
      if err != nil {
         t.Fatal(err)
      }
      resp := func() *http.Response {
         resp, err := http.Get(paramount_plus)
         if err != nil {
            t.Fatal(err)
         }
         defer resp.Body.Close()
         return resp
      }()
      fmt.Println(relay, resp.Status)
      if resp.StatusCode == http.StatusOK {
         break
      }
   }
}

func TestRelays(t *testing.T) {
   var relays public_relays
   err := relays.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, relay := range relays.Countries {
      fmt.Printf("%+v\n", relay)
   }
}
func TestConnect(t *testing.T) {
   err := connect("mx")
   if err != nil {
      t.Fatal(err)
   }
   var connect connection
   err = connect.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", connect)
   err = disconnect()
   if err != nil {
      t.Fatal(err)
   }
   err = connect.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", connect)
}

type connection struct {
   Country string
}

func (c *connection) New() error {
   resp, err := http.Get("https://ipv4.am.i.mullvad.net/json")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(c)
}
