package mullvad

import (
   "fmt"
   "net/http"
   "testing"
)

const paramount_plus = "https://www.paramountplus.com/apps-api/v2.0/androidphone/video/cid/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ.json?at=ABAAAAAAAAAAAAAAAAAAAAAAU%2FEmq2DGAxNGYi71fZdi2mb6UTU2%2BelcD3bGEhs6bo4%3D"

func TestParamount(t *testing.T) {
   var relays PublicRelays
   err := relays.New()
   if err != nil {
      t.Fatal(err)
   }
   defer Disconnect()
   for relay := range relays.Seq("Canada") {
      err = Connect(relay)
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
