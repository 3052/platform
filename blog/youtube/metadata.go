package main

import (
   "41.neocities.org/protobuf"
   "encoding/base64"
   "fmt"
)

const VISITOR_PRIVACY_METADATA = "CgJVUxIEGgAgIQ=="

func main() {
   data, err := base64.StdEncoding.DecodeString(VISITOR_PRIVACY_METADATA)
   if err != nil {
      panic(err)
   }
   message := protobuf.Message{}
   err = message.Unmarshal(data)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%#v\n", message)
}

var _ = protobuf.Message{
   1: {protobuf.Bytes("US")},
   2: {
      protobuf.Message{
         3: {protobuf.Bytes{}},
         4: {protobuf.Varint(33)},
      },
   },
}
