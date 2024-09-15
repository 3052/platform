package instagram

import (
   "fmt"
   "reflect"
   "testing"
)

func TestSize(t *testing.T) {
   size := reflect.TypeOf(&struct{}{}).Size()
   for _, test := range size_tests {
      if reflect.TypeOf(test).Size() > size {
         fmt.Printf("*%T\n", test)
      } else {
         fmt.Printf("%T\n", test)
      }
   }
}

var size_tests = []any{
   Address{},
   MediaData{},
}

const test_address = "https://instagram.com/apolloniallewellyn/p/C-A8xdkCu2m"

func TestMedia(t *testing.T) {
   var web Address
   web.Set(test_address)
   media, err := web.Media()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}
