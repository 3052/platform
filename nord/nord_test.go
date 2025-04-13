package nord

import (
   "fmt"
   "os"
   "testing"
)

func TestRead(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   data, err := os.ReadFile(home + "/platform/nord/servers")
   if err != nil {
      t.Fatal(err)
   }
   var servers1 servers
   err = servers1.unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   for _, server1 := range servers1 {
      if server1.proxy_ssl() {
         fmt.Printf("%+v\n", server1)
      }
   }
}

func TestWrite(t *testing.T) {
   data, err := get_servers(0)
   if err != nil {
      t.Fatal(err)
   }
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   err = os.WriteFile(home + "/platform/nord/servers", data, os.ModePerm)
   if err != nil {
      t.Fatal(err)
   }
}
