package main

import (
   "154.pages.dev/protobuf"
   "bytes"
   "io"
   "net/http"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["Priority"] = []string{"u=0, i"}
   req.Header["User-Agent"] = []string{"com.google.android.youtube/19.33.35(Linux; U; Android 9; en_US; Android SDK built for x86 Build/PSR1.180720.012) gzip"}
   req.Header["X-Goog-Api-Format-Version"] = []string{"2"}
   req.Header["X-Goog-Visitor-Id"] = []string{"Cgt1TFh6dDNWSlJUMCipha-2BjIKCgJVUxIEGgAgTjoMCAEgoJa1uZbV8OVm"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "youtubei.googleapis.com"
   req.URL.Path = "/youtubei/v1/get_watch"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(req_body.Marshal()))
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

var req_body = protobuf.Message{
   1: {
      protobuf.Message{
         1: {
            protobuf.Message{
               12: {protobuf.Bytes("unknown")},
               13: {protobuf.Bytes("Android SDK built for x86")},
               16: {protobuf.Varint(3)},
               17: {protobuf.Bytes("19.33.35")},
               18: {protobuf.Bytes("Android")},
               19: {protobuf.Bytes("9")},
               21: {protobuf.Unknown{
                  protobuf.Bytes("en-US"),
                  protobuf.Message{
                     12: {protobuf.Fixed32(1398091118)},
                  },
               }},
               22: {protobuf.Bytes("US")},
               37: {protobuf.Varint(411)},
               38: {protobuf.Varint(683)},
               39: {protobuf.Fixed32(1076138569)},
               40: {protobuf.Fixed32(1082699659)},
               41: {protobuf.Varint(3)},
               46: {protobuf.Varint(1)},
               49: {protobuf.Unknown{
                  protobuf.Bytes("\xaa\x03\x06310260"),
                  protobuf.Message{
                     53: {protobuf.Bytes("310260")},
                  },
               }},
               50: {protobuf.Varint(220920022)},
               52: {protobuf.Varint(4)},
               55: {protobuf.Varint(411)},
               56: {protobuf.Varint(683)},
               61: {protobuf.Varint(3)},
               62: {
                  protobuf.Message{
                     3: {protobuf.Bytes("CLSFr7YGEhQxNzE2MTUxNTkwMTQxMjcyMTMyORi0ha-2BjIyQU9qRm94Ml9rYnc4b2xtQ0hXWXFkZGx3a0MydFNScVRHU2JISjNhOHdHS3ZvX0lVemc6MkFPakZveDJfa2J3OG9sbUNIV1lxZGRsd2tDMnRTUnFUR1NiSEozYTh3R0t2b19JVXpnQtwBQ0FNU29BRU5RSnZTcVFMMkJzWVM1aEtpaHdiMkF1OEczQkhBTDhvQ2lRNk5CWlFQd0FQWEFxc0dxdy1iQlpjRmxBcDVxUUtTQmVrQXFRRnU1d1NSQTFvS2pnUDBCWlVFdWdJU3dRVURGVkt0bHRZTWdSMkZoUVhudVFXNnI5SUxuRDY3RjZDQ0JaaG51dDhGdVVTcVVJWWxoRTdRS0lnenRnanVTWWtYdjNEdUZkOFAteHZ5TGFpUkJ1QVFtUXpCQkkwdjNRYTRLNTRmRXlTRkpzRXNsUVllbnhFSg%3D%3D")},
                     5: {protobuf.Bytes("CLSFr7YGEhQxMjMzODAwMDc0MzA1Mzk2NTI4MBi0ha-2BiiU5PwSKLn1_BIojqL9Eiil0P0SKJ6R_hIo3LX-EijIyv4SKN3o_hIot-r-Eijs7P4SKIj3_hIowIP_Eiiuj_8SKK-S_xIor5T_EiiWlf8SKJaj_xIo0aT_Eijapf8SKJa0_xIogLb_Eijltv8SKPq2_xIos7f_Eii8u_8SKNO8_xIonr3_Eiipvv8SKJ7B_xIon8P_Eiitw_8SKIjE_xIo1sT_EijjxP8SKKLF_xIotcX_EijCxf8SKMXF_xIooMb_Eiisxv8SKO_G_xIo1Mf_EijYx_8SMjJBT2pGb3gyX2tidzhvbG1DSFdZcWRkbHdrQzJ0U1JxVEdTYkhKM2E4d0dLdm9fSVV6ZzoyQU9qRm94Ml9rYnc4b2xtQ0hXWXFkZGx3a0MydFNScVRHU2JISjNhOHdHS3ZvX0lVemdCdENBTVNVZzBib3RmNkZiNF83aU9ZSWNvRmtFZUNEdmtGUnNrVVdic0dKTW9ETWhVMTNjX0NES2liQW9LNEFnUzRxQUdaeG80QWtndXFBTTNKelFLZTlRT1VGc3llQlpZb19RdkFCcW43QmE4UHFmOEcyQWs9")},
                  },
               },
               64: {protobuf.Varint(28)},
               65: {protobuf.Fixed32(1076363264)},
               67: {protobuf.Varint(18446744073709551256)},
               77: {protobuf.Fixed32(1065353216)},
               78: {protobuf.Varint(1)},
               80: {protobuf.Unknown{
                  protobuf.Bytes("America/Denver"),
                  protobuf.Message{
                     8:  {protobuf.Fixed64(4913252798083261805)},
                     12: {protobuf.Fixed32(1919252078)},
                  },
               }},
               84: {
                  protobuf.Message{
                     4: {
                        protobuf.Varint(12586821118910807144),
                        protobuf.Varint(1120696127435781411),
                        protobuf.Varint(14887395423971444360),
                        protobuf.Varint(1334456436917326417),
                        protobuf.Varint(2063440235820364138),
                        protobuf.Varint(16159142388888625839),
                        protobuf.Varint(2371129479081436754),
                     },
                  },
               },
               86: {protobuf.Unknown{
                  protobuf.Bytes("\x12\x04\b\x03\x12\x00\x18\x00"),
                  protobuf.Message{
                     2: {protobuf.Unknown{
                        protobuf.Bytes("\b\x03\x12\x00"),
                        protobuf.Message{
                           1: {protobuf.Varint(3)},
                           2: {protobuf.Bytes("")},
                        },
                     }},
                     3: {protobuf.Varint(0)},
                  },
               }},
               92: {protobuf.Bytes("ranchu;")},
               97: {protobuf.Unknown{
                  protobuf.Bytes("\b\x01\x10\x15"),
                  protobuf.Message{
                     1: {protobuf.Varint(1)},
                     2: {protobuf.Varint(21)},
                  },
               }},
               98: {protobuf.Bytes("Android")},
               100: {protobuf.Unknown{
                  protobuf.Bytes("\n\x05\b\x96\x0f\x18\x01"),
                  protobuf.Message{
                     1: {protobuf.Unknown{
                        protobuf.Bytes("\b\x96\x0f\x18\x01"),
                        protobuf.Message{
                           1: {protobuf.Varint(1942)},
                           3: {protobuf.Varint(1)},
                        },
                     }},
                  },
               }},
               102: {protobuf.Unknown{
                  protobuf.Bytes("\x10\x03\x18\x01"),
                  protobuf.Message{
                     2: {protobuf.Varint(3)},
                     3: {protobuf.Varint(1)},
                  },
               }},
            },
         },
         3: {protobuf.Unknown{
            protobuf.Bytes("8\x00x\x00"),
            protobuf.Message{
               7:  {protobuf.Varint(0)},
               15: {protobuf.Varint(0)},
            },
         }},
         5: {protobuf.Unknown{
            protobuf.Bytes("\x8a\x02\x00"),
            protobuf.Message{
               33: {protobuf.Bytes("")},
            },
         }},
         6: {protobuf.Unknown{
            protobuf.Bytes("\x12\"\"\x13\b\xb4\xbeˣ\xaa\x91\x88\x03\x15\xc0n\x12\x01\x1d\x11\xb7\x02\x162\bexternal\x9a\x01\x00"),
            protobuf.Message{
               2: {protobuf.Unknown{
                  protobuf.Bytes("\"\x13\b\xb4\xbeˣ\xaa\x91\x88\x03\x15\xc0n\x12\x01\x1d\x11\xb7\x02\x162\bexternal\x9a\x01\x00"),
                  protobuf.Message{
                     4: {protobuf.Unknown{
                        protobuf.Bytes("\b\xb4\xbeˣ\xaa\x91\x88\x03\x15\xc0n\x12\x01\x1d\x11\xb7\x02\x16"),
                        protobuf.Message{
                           1: {protobuf.Varint(1724629696831284)},
                           2: {protobuf.Fixed32(17985216)},
                           3: {protobuf.Fixed32(369276689)},
                        },
                     }},
                     6:  {protobuf.Bytes("external")},
                     19: {protobuf.Bytes("")},
                  },
               }},
            },
         }},
         9: {
            protobuf.Message{
               1: {
                  protobuf.Message{
                     1: {protobuf.Bytes("ms")},
                     2: {protobuf.Bytes("CoACbEt23KnV1VkHxR6tOrbJmAz_iwAjJZ7MK9Isvunf105feqMTi14gVPgvNuNdH1yskbmsm4D7qtj0lCaqN1ibPzBuqwIVhXZVJn2L5IWOi8506BxjLHJQENTReO2S3XiL8sLDFjqehxnDavMZUwIBodGlvabO5h4rweBCisNO2mAAVPLU6N5QTJHtCh2T_RA59Vp1R9WPdwOJeRe0Y2DnODr1hBSX1P9oz3DGHug5UaB7_wh1wwsTgQSSLmne8IXBPZLCuNzjNjIh6zWSS4C1PcRqwD4SKzLNDYOawT7aACbLVPDmakTIWy6M_IanH0v5ajhW8yFYU6mXoF4uRJ6qAQqAAqk5LG0YUUYDEgdL8uhue6YxSZ9sHNJZdZ29MBSaYZnrGqydeMoZ2cDQ_LdCuo86HYxu_8v6WDifmsD7rgXVtusprFod_dTBukjyv_oh-8bjS3QMMzCB6egiUr_n6iQhDCPLIkodxBeiugsVqGZVRXL5fNCr3c1nhC4If_GNKloXVef28i2w2YOXHJTG592SsmLCMhhg_ff_XPaWP6BbldDUT8U9LGw4wHof6xTgYWexbAKQbICtuL3LgO5o58WHBOa--rZOP6e0wtz7nIMxl-pCXtFzr9mSMzo-rMIFV82fNdq2apQFdGu56XaIT7kmZYfzTInhGSxfSyff9_-Hyy4SEHQEXg4ZIkunAwSR2eLyIMM")},
                  },
               },
            },
         },
      },
   },
   2: {
      protobuf.Message{
         2: {protobuf.Bytes("40wkJJXfwQ0")},
         3: {protobuf.Varint(0)},
         4: {
            protobuf.Message{
               1: {
                  protobuf.Message{
                     3:  {protobuf.Bytes("android-google")},
                     4:  {protobuf.Varint(0)},
                     5:  {protobuf.Varint(83)},
                     6:  {protobuf.Varint(0)},
                     7:  {protobuf.Varint(3)},
                     8:  {protobuf.Varint(0)},
                     10: {protobuf.Varint(0)},
                     11: {protobuf.Varint(0)},
                     12: {protobuf.Bytes("sdkv=a.19.33.35&output=xml_vast2")},
                     29: {protobuf.Varint(0)},
                     31: {protobuf.Unknown{
                        protobuf.Bytes(":\x00"),
                        protobuf.Message{
                           7: {protobuf.Bytes("")},
                        },
                     }},
                     37: {protobuf.Varint(0)},
                     38: {protobuf.Varint(0)},
                     41: {protobuf.Varint(0)},
                  },
               },
            },
         },
         5:  {protobuf.Varint(0)},
         8:  {protobuf.Varint(0)},
         15: {protobuf.Varint(0)},
         23: {protobuf.Bytes("NFwPNbPq7fMlIl7h")},
         26: {protobuf.Bytes("")},
         27: {protobuf.Unknown{
            protobuf.Bytes("\n\x06:\x04\b\x01\x10\x01"),
            protobuf.Message{
               1: {protobuf.Unknown{
                  protobuf.Bytes(":\x04\b\x01\x10\x01"),
                  protobuf.Message{
                     7: {protobuf.Unknown{
                        protobuf.Bytes("\b\x01\x10\x01"),
                        protobuf.Message{
                           1: {protobuf.Varint(1)},
                           2: {protobuf.Varint(1)},
                        },
                     }},
                  },
               }},
            },
         }},
         28: {protobuf.Unknown{
            protobuf.Bytes("\b\x00\x10\x00\x18\x00"),
            protobuf.Message{
               1: {protobuf.Varint(0)},
               2: {protobuf.Varint(0)},
               3: {protobuf.Varint(0)},
            },
         }},
      },
   },
   3: {protobuf.Unknown{
      protobuf.Bytes("\x12\v40wkJJXfwQ02\x00H\x00P\x00x\x00\xc0\x01\x00\xc8\x01\x00\xd0\x01\x00\xe0\x01\x01\xf0\x01\x00\xa2\x02\x02(K\xf8\x02\x00\x82\x03\x00"),
      protobuf.Message{
         2:  {protobuf.Bytes("40wkJJXfwQ0")},
         6:  {protobuf.Bytes("")},
         9:  {protobuf.Varint(0)},
         10: {protobuf.Varint(0)},
         15: {protobuf.Varint(0)},
         24: {protobuf.Varint(0)},
         25: {protobuf.Varint(0)},
         26: {protobuf.Varint(0)},
         28: {protobuf.Varint(1)},
         30: {protobuf.Varint(0)},
         36: {protobuf.Unknown{
            protobuf.Bytes("(K"),
            protobuf.Message{
               5: {protobuf.Varint(75)},
            },
         }},
         47: {protobuf.Varint(0)},
         48: {protobuf.Bytes("")},
      },
   }},
}
