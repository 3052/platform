package main

import (
   "bufio"
   "github.com/refraction-networking/utls"
   "golang.org/x/net/http2"
   "net"
   "net/http"
)

// resp, err := android{}.RoundTrip(&req)

type android struct{}

func (android) RoundTrip(req *http.Request) (*http.Response, error) {
   config := tls.Config{ServerName: req.URL.Host}
   dial_conn, err := net.Dial("tcp", req.URL.Host+":443")
   if err != nil {
      return nil, err
   }
   defer dial_conn.Close()

tls_conn := tls.UClient(
   dial_conn, &config,
   
   //tls.HelloAndroid_11_OkHttp,
   //tls.Hello360_11_0 ,
   //tls.Hello360_7_5  ,
   //tls.HelloChrome_100,         
   //tls.HelloChrome_102         ,
   //tls.HelloChrome_106_Shuffle ,
   //tls.HelloChrome_115_PQ     ,
   //tls.HelloChrome_120 ,
   //tls.HelloChrome_120_PQ ,
   //tls.HelloChrome_58    ,      
   //tls.HelloChrome_62      ,    
   //tls.HelloChrome_70       ,   
   //tls.HelloChrome_72        ,  
   //tls.HelloChrome_83         , 
   //tls.HelloChrome_87          ,
   //tls.HelloChrome_96          ,
   //tls.HelloEdge_106  ,
   //tls.HelloEdge_85   ,
   //tls.HelloFirefox_102,  
   //tls.HelloFirefox_105 , 
   //tls.HelloFirefox_120  ,
   //tls.HelloFirefox_55   ,
   //tls.HelloFirefox_56   ,
   //tls.HelloFirefox_63   ,
   //tls.HelloFirefox_65   ,
   //tls.HelloFirefox_99   ,
   //tls.HelloIOS_11_1 ,
   //tls.HelloIOS_12_1 ,
   //tls.HelloIOS_13   ,
   //tls.HelloIOS_14   ,
   //tls.HelloQQ_11_1 ,
   tls.HelloSafari_16_0 ,
)
   
var h2 bool = true
//var h2 bool
   
   defer tls_conn.Close()
   
   if h2 {
         tr := http2.Transport{}
          cConn, err := tr.NewClientConn(tls_conn)
          if err != nil {
                  return nil, err
          }
          return cConn.RoundTrip(req)
   }
    err = req.Write(tls_conn)
    if err != nil {
            return nil, err
    }
    return http.ReadResponse(bufio.NewReader(tls_conn), req)
   
   
}
