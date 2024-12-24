package mullvad

import (
   "encoding/json"
   "fmt"
   "net/http"
   "testing"
)

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
