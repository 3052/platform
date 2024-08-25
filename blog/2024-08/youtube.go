package main

import (
   "154.pages.dev/protobuf"
   "io"
   "net/http"
   "net/url"
   "os"
   "bytes"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["User-Agent"] = []string{"com.google.android.youtube/18.18.38(Linux; U; Android 12; en_US; sdk_gphone64_x86_64 Build/SE1B.240122.005) gzip"}
   req.Header["X-Goog-Api-Format-Version"] = []string{"2"}
   req.Header["X-Goog-Visitor-Id"] = []string{"CgtzZjI3dkR5Z0VSVSjGla22BjIKCgJVUxIEGgAgYzoMCAEgtO6MnOXY0uVm"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "youtubei.googleapis.com"
   req.URL.Path = "/youtubei/v1/player"
   value := url.Values{}
   value["id"] = []string{"40wkJJXfwQ0"}
   value["key"] = []string{"AIzaSyA8eiZmM1FaDVjRy-df2KTyQ_vz_yYM39w"}
   value["t"] = []string{"pY_UR-mBysO5"}
   req.URL.RawQuery = value.Encode()
   req.URL.Scheme = "https"
   //req.Body = io.NopCloser(body)
   req.Body = io.NopCloser(bytes.NewReader(message.Marshal()))
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

var message = protobuf.Message{
   1: {protobuf.Message{
         1: {
            protobuf.Message{
               12: {protobuf.Bytes("Google")},
               13: {protobuf.Bytes("sdk_gphone64_x86_64")},
               16: {protobuf.Varint(3)},
               17: {protobuf.Bytes("18.18.38")},
               18: {protobuf.Bytes("Android")},
               19: {protobuf.Bytes("12")},
               21: {
                  protobuf.Bytes("en-US"),
               },
               22: {protobuf.Bytes("US")},
               25: {protobuf.Bytes("bbcbee93423ac1d5")},
               37: {protobuf.Varint(411)},
               38: {protobuf.Varint(715)},
               39: {protobuf.Fixed32(1076138569)},
               40: {protobuf.Fixed32(1083119089)},
               41: {protobuf.Varint(3)},
               46: {protobuf.Varint(1)},
               50: {protobuf.Varint(225014047)},
               52: {protobuf.Varint(4)},
               55: {protobuf.Varint(411)},
               56: {protobuf.Varint(715)},
               61: {protobuf.Varint(3)},
               62: {
                  protobuf.Message{
                     3: {protobuf.Bytes("CNKVrbYGEhQxNDY3MDY3MzQzMzI1MjAzNTEyOBjSla22BjIyQU9qRm94MktlbzhUc1hsckVORjFGT2VVek1Ic2NGVFpKeGlqcm8xN3hJT2hoN0xDM1E6MkFPakZveDJLZW84VHNYbHJFTkYxRk9lVXpNSHNjRlRaSnhpanJvMTd4SU9oaDdMQzNRQmhDQU1TU2cwbm05S3BBdllHeGhMbUV2WExCcDRNMnd2N0FjNE4xUkswRjJfbUJKRURyUUlWSWEyVzFneUJIWVdGQmVlNUJkejAzQXVvZnR0X3gwVF9XZi1vQnFMR0JvSWt3QVRpTHY0cA%3D%3D")},
                     5: {protobuf.Bytes("CNKVrbYGEhQxMDIwOTUxNDkwNDQ0Mjg2OTc4MxjSla22BiiU5PwSKLn1_BIojqL9Eiil0P0SKJ6R_hIo3ej-Eii36v4SKOft_hIowIP_Eiivj_8SKLCS_xIotKP_EiiAtv8SKNO8_xIonsH_Eiifw_8SKNfE_xIo48T_Eiiixf8SKLXF_xIoxMX_EjIyQU9qRm94MktlbzhUc1hsckVORjFGT2VVek1Ic2NGVFpKeGlqcm8xN3hJT2hoN0xDM1E6MkFPakZveDJLZW84VHNYbHJFTkYxRk9lVXpNSHNjRlRaSnhpanJvMTd4SU9oaDdMQzNRQjRDQU1TSWcwS290ZjZGYjRfNklvRzBpWEtBeFVXM2NfQ0RLaWJBcGFENGd1TDZnS3Bfd1k9")},
                  },
               },
               64: {protobuf.Varint(32)},
               65: {protobuf.Fixed32(1076363264)},
               67: {protobuf.Varint(18446744073709551316)},
               78: {protobuf.Varint(1)},
               80: {protobuf.Bytes("America/Chicago")},
               84: {
                  protobuf.Message{
                     4: {
                        protobuf.Varint(14887395423971444360),
                        protobuf.Varint(2371129479081436754),
                        protobuf.Varint(16159142388888625839),
                     },
                  },
               },
               86: {
                  protobuf.Message{
                     2: {
                        protobuf.Message{
                           1: {protobuf.Varint(3)},
                           2: {protobuf.Bytes("")},
                        },
                     },
                     3: {protobuf.Varint(0)},
                  },
               },
               92: {
                  protobuf.Bytes("AOSP;ranchu"),
               },
               97: {
                  protobuf.Message{
                     1: {protobuf.Varint(1)},
                     2: {protobuf.Varint(17)},
                  },
               },
               98: {protobuf.Bytes("google")},
               100: {
                  protobuf.Message{
                     1: {
                        protobuf.Message{
                           1: {protobuf.Varint(1778)},
                           3: {protobuf.Varint(1)},
                        },
                     },
                  },
               },
               102: {
                  protobuf.Message{
                     1: {protobuf.Bytes("Unknown Renderer")},
                     2: {protobuf.Varint(3)},
                     3: {protobuf.Varint(1)},
                  },
               },
            },
         },
         3: {
            protobuf.Message{
               7:  {protobuf.Varint(0)},
               15: {protobuf.Varint(0)},
            },
         },
         6: {
            protobuf.Message{
               2: {
                  protobuf.Message{
                     4: {
                        protobuf.Message{
                           1: {protobuf.Varint(1724599003173475)},
                           2: {protobuf.Fixed32(17982453)},
                           3: {protobuf.Fixed32(404995197)},
                        },
                     },
                     6:  {protobuf.Bytes("external")},
                     12: {protobuf.Bytes("com.android.shell")},
                  },
               },
            },
         },
         9: {
            protobuf.Message{
               1: {
                  protobuf.Message{
                     1: {protobuf.Bytes("ms")},
                     2: {protobuf.Bytes("CoACV5JU_QJ2zna1U_EfEasYVMKYCr9gnk_wbDyu2zuPVYU0bhUNHXHEiM4GawcGrwQhvJRWbzh5S2XrK158ZYh41-WT01r-aK_nUuDxzki550BCZvMLNp-i4URoOLT81V6_hP2e41hQfwu5e2L_WmRqRFKdn7eQOK8_zj8o423DAKzwQ98b7w3JzNnvCgBX6FtBADTakmlUWk2eAITPj0Vgd7eTKIrZwlTbdKhOl7DsGnfwderWmVqZO933Pc5wsK3HsCO8d3Vf1t1ydI-68yjtK7nx1ebiJVID2Y12ERWMC_ko0snseV_ltkirt0j1BNmvkhzXmoI83HO-bD4Y7V2ePAqAAqJOFPYIYYAfCdaHR8E2CQZQG40WfThG2Fvvuff7vteDrPHnCQSCnKZnYwUGG04GkyQl_wv56_coXT_JTTIeVAIxSqcjwhLG3_ZhRHamI_KldzYLywUi7LPeY66hQkjZf0XwZoImGCsGQwKrTnAR48faSnXdUYh-dA2xz092-oMPAlXmm7zho_PaD6Ing8obhuUo2gaadvls5hjvgjPLBL1mkBEHBG-KmR3U7Dx_FHhHq6Br4NArH6QGcp1OLavHt7g4OIuh4MBxf6RAO3PCJcsOKfQVugwoHM-OjVGZyeiPVLEMllJJTZl4cpn6uXgilHIlcJHhl2gU0qI0Qgx39hISEITjbFQ6HZB5PhmvjJG3Y04")},
                  },
               },
            },
         },
      },
   },
   2: {protobuf.Bytes("40wkJJXfwQ0")},
   3: {protobuf.Varint(0)},
   4: {
      protobuf.Message{
         1: {
            protobuf.Message{
               3:  {protobuf.Bytes("mvapp-unknown")},
               4:  {protobuf.Varint(0)},
               5:  {protobuf.Varint(203)},
               6:  {protobuf.Varint(0)},
               7:  {protobuf.Varint(3)},
               8:  {protobuf.Varint(0)},
               10: {protobuf.Varint(0)},
               11: {protobuf.Varint(0)},
               12: {protobuf.Bytes("sdkv=a.18.18.38&output=xml_vast2")},
               29: {protobuf.Varint(0)},
               31: {
                  protobuf.Message{
                     7: {protobuf.Bytes("")},
                  },
               },
               37: {protobuf.Varint(0)},
               38: {protobuf.Varint(0)},
               41: {protobuf.Varint(0)},
            },
         },
      },
   },
   5:  {protobuf.Varint(0)},
   8:  {protobuf.Varint(0)},
   12: {protobuf.Bytes("6AQB")},
   23: {protobuf.Bytes("hLZdY4MSdi4zlh-2")},
   26: {protobuf.Bytes("")},
   28: {
      protobuf.Message{
         1: {protobuf.Varint(0)},
         2: {protobuf.Varint(0)},
         3: {protobuf.Varint(0)},
      },
   },
}
