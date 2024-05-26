package soundcloud

import (
   "encoding/json"
   "os"
   "testing"
)

func TestNextData(t *testing.T) {
   var next next_data
   err := next.New()
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(next)
}
