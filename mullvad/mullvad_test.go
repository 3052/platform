package mullvad

import (
   "errors"
   "log"
   "net/http"
   "testing"
)

var tests = []string{
   "ca-tor-ovpn-101",
   "ca-yyc-wg-201",
}

func Test(t *testing.T) {
   log.SetFlags(log.Ltime)
   for _, test1 := range tests{
      err := head(test1)
      if err != nil {
         t.Fatal(err)
      }
   }
}

const paramount = "https://www.paramountplus.com/apps-api/v2.0/androidphone/video/cid/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ.json?at=ABAAAAAAAAAAAAAAAAAAAAAAU%2FEmq2DGAxNGYi71fZdi2mb6UTU2%2BelcD3bGEhs6bo4%3D"

func head(server string) error {
   defer Disconnect()
   err := Connect(server)
   if err != nil {
      return err
   }
   log.Printf("%v", paramount)
   resp, err := http.Head(paramount)
   if err != nil {
      return err
   }
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   return nil
}
