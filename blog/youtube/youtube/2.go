package main

import (
   "154.pages.dev/protobuf"
   "fmt"
   "os"
)

func main() {
   data, err := os.ReadFile("response.txt")
   if err != nil {
      panic(err)
   }
   message := protobuf.Message{}
   err = message.Unmarshal(data)
   if err != nil {
      panic(err)
   }
   message, _ = message.Get(1)()
   message, _ = message.Get(2)()
   message, _ = message.Get(4)()
   message, _ = message.Get(3)()
   address, _ := message.GetBytes(2)()
   fmt.Println(string(address))
}
