package instagram

import (
   "fmt"
   "reflect"
   "testing"
)

func TestAddress(t *testing.T) {
   var web address
   web.Set("https://instagram.com/apolloniallewellyn/p/C-A8xdkCu2m")
   fmt.Printf("%+v\n", web)
}

func TestSize(t *testing.T) {
   size := reflect.TypeOf(&struct{}{}).Size()
   var test address
   if reflect.TypeOf(test).Size() > size {
      fmt.Printf("*%T\n", test)
   } else {
      fmt.Printf("%T\n", test)
   }
}

func TestMedia(t *testing.T) {
   var media media_data
   err := media.New("C-A8xdkCu2m")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}
