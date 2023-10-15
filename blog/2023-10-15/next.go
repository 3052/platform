package main

import (
   "154.pages.dev/protobuf"
   "bytes"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["User-Agent"] = []string{"com.google.android.youtube/16.49.39(Linux; U; Android 12; en_US; sdk_gphone64_x86_64 Build/SE1B.220616.007) gzip"}
   req.Header["X-Goog-Api-Format-Version"] = []string{"2"}
   req.Header["X-Goog-Visitor-Id"] = []string{"Cgt6d19oSjhzQktCayjEta-pBjIICgJVUxICGgA6CiCvt-qkxdj2lWU%3D"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "youtubei.googleapis.com"
   req.URL.Path = "/youtubei/v1/next"
   val := make(url.Values)
   val["key"] = []string{"AIzaSyA8eiZmM1FaDVjRy-df2KTyQ_vz_yYM39w"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(req_body.Append(nil)))
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}

var req_body = protobuf.Message{
   protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 12, Type: 2,  Value: protobuf.Bytes("Google")},
         protobuf.Field{Number: 13, Type: 2,  Value: protobuf.Bytes("sdk_gphone64_x86_64")},
         protobuf.Field{Number: 16, Type: 0,  Value: protobuf.Varint(3)},
         protobuf.Field{Number: 17, Type: 2,  Value: protobuf.Bytes("16.49.39")},
         protobuf.Field{Number: 18, Type: 2,  Value: protobuf.Bytes("Android")},
         protobuf.Field{Number: 19, Type: 2,  Value: protobuf.Bytes("12")},
         protobuf.Field{Number: 21, Type: 2,  Value: protobuf.Bytes("en-US")},
         protobuf.Field{Number: 22, Type: 2,  Value: protobuf.Bytes("US")},
         protobuf.Field{Number: 37, Type: 0,  Value: protobuf.Varint(432)},
         protobuf.Field{Number: 38, Type: 0,  Value: protobuf.Varint(848)},
         protobuf.Field{Number: 39, Type: 5,  Value: protobuf.Fixed32(1076677837)},
         protobuf.Field{Number: 40, Type: 5,  Value: protobuf.Fixed32(1084856730)},
         protobuf.Field{Number: 41, Type: 0,  Value: protobuf.Varint(3)},
         protobuf.Field{Number: 46, Type: 0,  Value: protobuf.Varint(1)},
         protobuf.Field{Number: 50, Type: 0,  Value: protobuf.Varint(225014047)},
         protobuf.Field{Number: 52, Type: 0,  Value: protobuf.Varint(4)},
         protobuf.Field{Number: 55, Type: 0,  Value: protobuf.Varint(432)},
         protobuf.Field{Number: 56, Type: 0,  Value: protobuf.Varint(848)},
         protobuf.Field{Number: 61, Type: 0,  Value: protobuf.Varint(3)},
         protobuf.Field{Number: 62, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 3, Type: 2,  Value: protobuf.Bytes("CMa1r6kGEhQxNzc0MTMzOTYwNTc5MzAxNDYxOBjGta-pBjIyQU9qRm94MkoxMmdGdDd5N3FJankyRUh1ZmZmek5EbmdsTmVELXY4cGxkR3dqUVBvbUE6MkFPakZveDJKMTJnRnQ3eTdxSWp5MkVIdWZmZnpORG5nbE5lRC12OHBsZEd3alFQb21BQkRDQU1TTHcwVjh0MnBBb2d4eUJiLUJmQXFfUWFQR2NrSUZSN0E4OVlNeXBjQzJfWUdqaDZ5UnQ1aW5TLURoQVdzdHNvTA%3D%3D")},
            protobuf.Field{Number: 5, Type: 2,  Value: protobuf.Bytes("CMa1r6kGEhQxODA4NTEzNjQzNDM3NjM3MjM3OBjGta-pBiiU5PwSKLn1_BIo25P9EiiOov0SKMay_RIoqrT9Eiiekf4SKJqt_hIoiMj-EijIyv4SKN3O_hIoqOH-Eij95v4SKLfq_hIohuz-Eiig8f4SKNTz_hIokPf-EjIyQU9qRm94MkoxMmdGdDd5N3FJankyRUh1ZmZmek5EbmdsTmVELXY4cGxkR3dqUVBvbUE6MkFPakZveDJKMTJnRnQ3eTdxSWp5MkVIdWZmZnpORG5nbE5lRC12OHBsZEd3alFQb21BQjRDQU1TSVEwTDJJXzVGY29BM1REVEFjQVJGUmFONHMwTWktNEI4ZUFPdFlBRWc2X05Ddz09")},
         }},
         protobuf.Field{Number: 64, Type: 0,  Value: protobuf.Varint(32)},
         protobuf.Field{Number: 65, Type: 5,  Value: protobuf.Fixed32(1075838976)},
         protobuf.Field{Number: 67, Type: 0,  Value: protobuf.Varint(18446744073709551316)},
         protobuf.Field{Number: 78, Type: 0,  Value: protobuf.Varint(1)},
         protobuf.Field{Number: 80, Type: 2,  Value: protobuf.Bytes("America/Chicago")},
         protobuf.Field{Number: 84, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 4, Type: 0,  Value: protobuf.Varint(7030543082742800693)},
            protobuf.Field{Number: 4, Type: 0,  Value: protobuf.Varint(14432795415577216842)},
            protobuf.Field{Number: 4, Type: 0,  Value: protobuf.Varint(7955987839527910596)},
            protobuf.Field{Number: 4, Type: 0,  Value: protobuf.Varint(15701289610270850466)},
            protobuf.Field{Number: 4, Type: 0,  Value: protobuf.Varint(11151738735368279274)},
         }},
         protobuf.Field{Number: 89, Type: 2,  Value: protobuf.Bytes("")},
         protobuf.Field{Number: 97, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(1)},
         }},
         protobuf.Field{Number: 98, Type: 2,  Value: protobuf.Bytes("google")},
      }},
      protobuf.Field{Number: 3, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 7, Type: 0,  Value: protobuf.Varint(0)},
         protobuf.Field{Number: 15, Type: 0,  Value: protobuf.Varint(0)},
      }},
      protobuf.Field{Number: 6, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 2, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 4, Type: 2, Value: protobuf.Prefix{
               protobuf.Field{Number: 1, Type: 0,  Value: protobuf.Varint(1697372883439256)},
               protobuf.Field{Number: 2, Type: 5,  Value: protobuf.Fixed32(132489275)},
               protobuf.Field{Number: 3, Type: 5,  Value: protobuf.Fixed32(402734815)},
            }},
            protobuf.Field{Number: 6, Type: 2,  Value: protobuf.Bytes("external")},
            protobuf.Field{Number: 12, Type: 2,  Value: protobuf.Bytes("com.android.shell")},
         }},
      }},
      protobuf.Field{Number: 9, Type: 2, Value: protobuf.Prefix{
         protobuf.Field{Number: 1, Type: 2, Value: protobuf.Prefix{
            protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("ms")},
            protobuf.Field{Number: 2, Type: 2,  Value: protobuf.Bytes("CoACNUSbWEkoOeMkQbU_L51GMRq3E-_OVsP_aEliM6AJJs_MdaN9WFF_J5giTymq2jgDps9PSITZ6ENjuCUtJP1QuxvCNRYt0KyVlwVjdLk6fkiTDKXKQoYuCallJvLLycqGRZFuDzicQijb_5ggaWWQrlWX0NdAm3hWTrnzCa68vXNEiZIzBFwPafsYi7RB5lCNHQHetOp14LVl4fUGoLrNwwIDAHzEQ64vsG8Wn0wmUfA4Y5Hw1bXR_L8pITF6FI1awx60o6NISSeLxnDMw1lFTlAHOfmrVkesKohnZoV2pz4o_Xwtd4TXzFD16YH8mF9fzgUmLkU3zw3HQWSe8fmabBIQZGaRsYEPfv0b_tvxlbX6JQ")},
         }},
      }},
   }},
   protobuf.Field{Number: 2, Type: 2,  Value: protobuf.Bytes("2ZcDwdXEVyI")},
   protobuf.Field{Number: 6, Type: 2,  Value: protobuf.Bytes("")},
   protobuf.Field{Number: 9, Type: 0,  Value: protobuf.Varint(0)},
   protobuf.Field{Number: 10, Type: 0,  Value: protobuf.Varint(0)},
   protobuf.Field{Number: 15, Type: 0,  Value: protobuf.Varint(0)},
   protobuf.Field{Number: 24, Type: 0,  Value: protobuf.Varint(0)},
   protobuf.Field{Number: 25, Type: 0,  Value: protobuf.Varint(0)},
   protobuf.Field{Number: 26, Type: 0,  Value: protobuf.Varint(0)},
   protobuf.Field{Number: 28, Type: 0,  Value: protobuf.Varint(1)},
   protobuf.Field{Number: 30, Type: 0,  Value: protobuf.Varint(0)},
   protobuf.Field{Number: 36, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 5, Type: 0,  Value: protobuf.Varint(561)},
   }},
   protobuf.Field{Number: 47, Type: 0,  Value: protobuf.Varint(0)},
   protobuf.Field{Number: 48, Type: 2,  Value: protobuf.Bytes("")},
   protobuf.Field{Number: 51, Type: 2, Value: protobuf.Prefix{
      protobuf.Field{Number: 1, Type: 2,  Value: protobuf.Bytes("com.android.chrome")},
      protobuf.Field{Number: 2, Type: 0,  Value: protobuf.Varint(463807437)},
   }},
}
