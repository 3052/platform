package youtube

import (
   "fmt"
   "testing"
   "time"
)

func TestCode(t *testing.T) {
   var code DeviceCode
   err := code.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf(
      "1. go to\n%v\n\n2. enter this code\n%v\n",
      code.VerificationUrl, code.UserCode,
   )
   for range 9 {
      time.Sleep(9 * time.Second)
      auth, err := code.Auth()
      if err != nil {
         t.Fatal(err)
      }
      if err := auth.Unmarshal(); err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", auth)
      if auth.V.AccessToken != "" {
         break
      }
   }
}
