package youtube

import (
   "fmt"
   "testing"
)

func TestCodeWrite(t *testing.T) {
   var code DeviceCode
   err := code.New()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.json", code.Data, 0666)
   err = code.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(code)
}

func TestCodeRead(t *testing.T) {
   var (
      code DeviceCode
      err error
   )
   code.Data, err = os.ReadFile("code.json")
   if err != nil {
      t.Fatal(err)
   }
   err = code.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   token, err := code.Token()
   if err != nil {
      t.Fatal(err)
   }
   err = token.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", token)
}
