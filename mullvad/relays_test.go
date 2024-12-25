package mullvad

import (
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
   for relay := range relays.relays("Canada") {
      fmt.Println(relay)
      code := func() int {
         err = connect(relay)
         if err != nil {
            t.Fatal(err)
         }
         defer disconnect()
         resp, err := http.Get(paramount_plus)
         if err != nil {
            t.Fatal(err)
         }
         defer resp.Body.Close()
         return resp.StatusCode
      }()
      if code == http.StatusOK {
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
