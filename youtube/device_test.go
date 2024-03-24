package youtube

import (
   "fmt"
   "testing"
   "time"
)

func TestCode(t *testing.T) {
   var code DeviceCode
   err := code.Post()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf(
      "1. go to\n%v\n\n2. enter this code\n%v\n",
      code.Verification_URL, code.User_Code,
   )
   for range [9]bool{} {
      time.Sleep(9 * time.Second)
      auth, err := code.OAuth()
      if err != nil {
         t.Fatal(err)
      }
      if err := auth.Unmarshal(); err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", auth)
      if auth.V.Access_Token != "" {
         break
      }
   }
}
