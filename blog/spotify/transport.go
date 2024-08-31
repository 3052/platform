package main

import (
   "bufio"
   "encoding/base64"
   "github.com/refraction-networking/utls"
   "golang.org/x/net/http2"
   "net"
   "net/http"
)

const spotify = "FgMDAKQBAACgAwP16VkGJcLhMqzoTxXG4W+IswTmOXGyFs7vSCL3oOd9CwAAFsArwCzAL8AwwBPAFACcAJ0ALwA1AP8BAABhAAAAFwAVAAASbG9naW41LnNwb3RpZnkuY29tABcAAAAjAAAADQAWABQGAQYDBQEFAwQBBAMDAQMDAgECAwAQAA4ADAJoMghodHRwLzEuMQALAAIBAAAKAAgABgAXABgAGQ=="

var h2 bool = true

type android struct{}

func (android) RoundTrip(req *http.Request) (*http.Response, error) {
   config := tls.Config{ServerName: req.URL.Host}
   dial_conn, err := net.Dial("tcp", req.URL.Host+":443")
   if err != nil {
      return nil, err
   }
   defer dial_conn.Close()
   tls_conn := tls.UClient(dial_conn, &config, tls.HelloCustom)
   defer tls_conn.Close()
   data, err := base64.StdEncoding.DecodeString(spotify)
   if err != nil {
      return nil, err
   }
   var finger tls.Fingerprinter
   spec, err := finger.FingerprintClientHello(data)
   if err != nil {
      return nil, err
   }
   err = tls_conn.ApplyPreset(spec)
   if err != nil {
      return nil, err
   }
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
