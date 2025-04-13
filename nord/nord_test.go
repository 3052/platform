package nord

import (
   "fmt"
   "testing"
)

func Test(t *testing.T) {
   var servers1 servers
   err := servers1.New(0)
   if err != nil {
      t.Fatal(err)
   }
   for _, server := range servers1.Servers {
      fmt.Printf("%+v\n", server)
   }
}
