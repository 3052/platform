package youtube

import (
   "fmt"
   "os"
   "testing"
)

func TestCodeRead(t *testing.T) {
   text, err := os.ReadFile("code.txt")
   if err != nil {
      t.Fatal(err)
   }
   var code DeviceCode
   err = code.Unmarshal(text)
   if err != nil {
      t.Fatal(err)
   }
   token, err := code.Token()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", token)
}

func TestCodeWrite(t *testing.T) {
   var code DeviceCode
   err := code.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(code)
   text, err := code.Marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.txt", text, 0666)
}
