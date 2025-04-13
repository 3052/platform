package nord

import (
   "fmt"
   "os"
   "testing"
)

func TestServerLoad(t *testing.T) {
   load := server_loads{
      {Count:2, Country:"CA", City:"toronto", Hostname:"ca1788.nordvpn.com"},
      {Count:3, Country:"CA", City:"toronto", Hostname:"ca1789.nordvpn.com"},
   }
   for range 9 {
      fmt.Println(load.country("CA"))
   }
}

func TestCountry(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   data, err := os.ReadFile(home + "/platform/nord/server_loads")
   if err != nil {
      t.Fatal(err)
   }
   var loads server_loads
   err = loads.unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   for _, load := range loads {
      if load.Country == "CA" {
         fmt.Printf("%+v\n\n", load)
      }
   }
}

func TestWrite(t *testing.T) {
   servers, err := get_servers(0)
   if err != nil {
      t.Fatal(err)
   }
   data, err := get_server_loads(servers).marshal()
   if err != nil {
      t.Fatal(err)
   }
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   err = os.WriteFile(home+"/platform/nord/server_loads", data, os.ModePerm)
   if err != nil {
      t.Fatal(err)
   }
}
