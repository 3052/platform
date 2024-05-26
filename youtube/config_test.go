package youtube

import (
   "encoding/json"
   "os"
   "testing"
)

func TestConfig(t *testing.T) {
   var tube innertube
   err := tube.New()
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(tube)
}
