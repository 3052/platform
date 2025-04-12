package nord

import (
   "fmt"
   "testing"
)

func TestSocks(t *testing.T) {
   servers1, err := servers(9999)
   if err != nil {
      t.Fatal(err)
   }
   for _, server1 := range servers1 {
      for _, technology1 := range server1.Technologies {
         if technology1.Identifier == "socks" {
            fmt.Printf("%+v\n", server1)
         }
      }
   }
}

func TestTechnology(t *testing.T) {
   servers1, err := servers(9999)
   if err != nil {
      t.Fatal(err)
   }
   technologies := map[technology]struct{}{}
   for _, server1 := range servers1 {
      for _, technology1 := range server1.Technologies {
         technologies[technology1] = struct{}{}
      }
   }
   for technology1 := range technologies {
      fmt.Printf("%+v\n", technology1)
   }
}
