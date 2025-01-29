package webDriver

import (
   "os"
   "testing"
)

func TestNavigate(t *testing.T) {
   var session session_state
   err := session.New()
   if err != nil {
      t.Fatal(err)
   }
   resp, err := session.navigate("https://www.sky.ch")
   if err != nil {
      t.Fatal(err)
   }
   resp.Write(os.Stdout)
}
