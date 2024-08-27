package main

import (
   "154.pages.dev/protobuf"
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "time"
)

func main() {
   var req http.Request
   req.URL = &url.URL{}
   req.URL.Path = "/youtubei/v1/get_watch"
   req.Header = http.Header{}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL.Host = "youtubei.googleapis.com"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(Request.Marshal()))
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["X-Goog-Visitor-Id"] = []string{"Cgt0OV94ZHc0Q1hnQSjUja-2BjIKCgJVUxIEGgAgTzoMCAEgwKfTncLa8eVm"}
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      panic(resp.Status)
   }
   data, err := io.ReadAll(resp.Body)
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
   address := func() string {
      b, _ := message.GetBytes(2)()
      return string(b)
   }()
   fmt.Println(address)
   after := time.After(30 * time.Second)
   for {
      select {
      case <-time.After(time.Second):
         resp, err := http.Head(address)
         if err != nil {
            panic(err)
         }
         fmt.Println(resp.Status)
      case <-after:
         return
      }
   }
}

var Request = protobuf.Message{
   1: {protobuf.Message{ // keep
         1: {protobuf.Message{ // keep
               16: {protobuf.Varint(3)},
               17: {protobuf.Bytes("19.33.35")},
               19: {protobuf.Bytes("9")},
               64: {protobuf.Varint(28)},
            },
         },
      },
   },
   2: { protobuf.Message{ // keep
         2: {protobuf.Bytes("40wkJJXfwQ0")}, // keep
         27: { protobuf.Message{
               1: { protobuf.Message{
                     1: { protobuf.Message{
                           1: {protobuf.Bytes("\x01/~e\x86\b\x11\xe8\xd4\xd1\x0fh\x7fm\xef\xfc\xb4n\t\xc6\xeaI\xa5,6\x99\xb1\xb6\xe8\xc4ܷ\x94\x90\x92\xa9K\x00\x1d\xfa?\xb6\x8e\x8d`\xfc7E\x90\x88\x9d\xfd\xcbˁ\xc5\xccnI\x16b\x03s\xb1om\xdd\xf6\xa9\xe4\xc8V\x19i\xab\x81\xd2\x19kZ\xd5\xfc\x11\xed\xaa f\x8e\xb5\xcc\x02\x89\xf8j\x11\xfe\x06m\a\x8dF\xe7/tX\xa6\xebȰ\xf8\xaa\x83\xa7ʜ\xd3m\xcbҡC\xda\xc3m\xf5\xad\xa4\xa0\xbc\x16bA\xd6\xc9\x05\xb5\xbe\xb7o\x99\xbb\v\x1e\xf8\xb3\xc8iP\xbcQ\xe7\xa9Qo\xa0O\xc6i\u17ef|?!\x89\x0e+\xff\x8fS\xce\x11\xbbߚ\xe6\xb4\xd75\xe2U\xdf\xf3\xb4\xbf\xd9;\xd0̝\xdd鎫K\xc3\xff\xee\xf3H\x1b\xd5*u\xa3;\x03\xe9\x13η\xc2\xeb\x8fG\xb7}\x16\x04J\xa6aR˜\x97.\xdf\\r\t\xbdIẻ\xd1\xf6\x9a{\n\x9e˴Ռ\x8d!\xecH")},
                           2: {protobuf.Bytes("\x01\x027\xfd\xb6\x1a\x91\xc5\xda\x19\xf2`k\x9f\xda\xe0\x18Bk\x00Y\xc4:\xcbٶI\xed7\xf9\xfd2\x1ae\x9c\x9a^\x91\xb8\xa2N\x0eԚڐ\x06\xbe\x81\x9a\xeff\x9a")},
                        },
                     },
                  },
               },
            },
         },
      },
   },
}
