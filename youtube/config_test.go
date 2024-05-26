package youtube

import (
   "encoding/json"
   "os"
   "testing"
)

func TestConfig(t *testing.T) {
   var con config
   err := con.get()
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(con)
}
