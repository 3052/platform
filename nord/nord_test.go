package nord

import (
   "os"
   "testing"
)

func Test(t *testing.T) {
   resp, err := servers()
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   file, err := os.Create("nord.json")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   file.ReadFrom(resp.Body)
}
