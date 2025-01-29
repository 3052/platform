package webDriver

import (
   "fmt"
   "testing"
)

func TestCookie(t *testing.T) {
   var session session_state
   err := session.New()
   if err != nil {
      t.Fatal(err)
   }
   err = session.navigate("https://httpbingo.org/cookies/set?k1=v1&k2=v2")
   if err != nil {
      t.Fatal(err)
   }
   cookie, err := session.cookie()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", cookie)
}
