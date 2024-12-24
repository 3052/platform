package mullvad

import (
   "os"
   "testing"
)

func TestRelays(t *testing.T) {
   resp, err := relays()
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   file, err := os.Create("relays.json")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   file.ReadFrom(resp.Body)
}
