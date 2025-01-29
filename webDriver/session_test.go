package webDriver

import (
   "fmt"
   "testing"
)

func TestSession(t *testing.T) {
   var session session_state
   err := session.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", session)
}
