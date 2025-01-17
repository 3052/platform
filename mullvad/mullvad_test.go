package mullvad

import (
   "errors"
   "fmt"
   "net/http"
   "os"
   "testing"
)

func TestPath(t *testing.T) {
   if Connect("") == nil {
      t.Fatal("Connect")
   }
   defer os.Setenv("PATH", os.Getenv("PATH"))
   os.Setenv("PATH", "")
   if Disconnect() == nil {
      t.Fatal("Disconnect")
   }
}

func TestServerList(t *testing.T) {
   var list ServerList
   t.Run("OpenVpn", func(t *testing.T) {
      t.Run("0", func(t *testing.T) {
         http.DefaultClient.Transport = transport{}
         if list.OpenVpn() == nil {
            t.Fatal(list)
         }
         http.DefaultClient.Transport = nil
      })
      t.Run("1", func(t *testing.T) {
         err := list.OpenVpn()
         if err != nil {
            t.Fatal(err)
         }
         err = list.head("ca-tor-ovpn-101")
         if err != nil {
            t.Fatal(err)
         }
      })
   })
   t.Run("Seq", func(t *testing.T) {
      err := list.OpenVpn()
      if err != nil {
         t.Fatal(err)
      }
      for range list.Seq("Canada") {
         break
      }
   })
   t.Run("WireGuard", func(t *testing.T) {
      t.Run("true", func(t *testing.T) {
         http.DefaultClient.Transport = transport{}
         if list.WireGuard() == nil {
            t.Fatal(list)
         }
         http.DefaultClient.Transport = nil
      })
      t.Run("false", func(t *testing.T) {
         err := list.WireGuard()
         if err != nil {
            t.Fatal(err)
         }
         err = list.head("ca-yyc-wg-201")
         if err != nil {
            t.Fatal(err)
         }
      })
   })
}

func (s ServerList) head(server string) error {
   defer Disconnect()
   fmt.Printf("%v Connect", server)
   err := Connect(server)
   if err != nil {
      return err
   }
   fmt.Println(" HEAD")
   resp, err := http.Head(paramount)
   if err != nil {
      return err
   }
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   return nil
}

func (transport) RoundTrip(*http.Request) (*http.Response, error) {
   return nil, errors.New("transport.RoundTrip")
}

type transport struct{}

const paramount = "https://www.paramountplus.com/apps-api/v2.0/androidphone/video/cid/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ.json?at=ABAAAAAAAAAAAAAAAAAAAAAAU%2FEmq2DGAxNGYi71fZdi2mb6UTU2%2BelcD3bGEhs6bo4%3D"
