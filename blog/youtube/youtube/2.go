package main

import (
   "blog/youtube"
   "fmt"
)

func main() {
   message, _ := youtube.Response.Get(4)()
   messages := message.Get(3)
   for {
      message, ok := messages()
      if !ok {
         break
      }
      if v, ok := message.GetVarint(1)(); ok {
         fmt.Println(v)
      }
      if v, ok := message.GetBytes(2)(); ok {
         fmt.Println(string(v))
      }
      fmt.Println()
   }
}
