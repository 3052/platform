package main

import (
   "41.neocities.org/platform/mullvad"
   "fmt"
   "net/http"
)

func (f *flags) get() (*http.Response, error) {
   resp, err := http.Get(f.address)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return resp, nil
}

func (f *flags) relay() error {
   http.DefaultClient.Transport = transport{}
   var relays mullvad.PublicRelays
   err := relays.New()
   if err != nil {
      return err
   }
   defer mullvad.Disconnect()
   for relay := range relays.Seq(f.location) {
      err = mullvad.Connect(relay)
      if err != nil {
         return err
      }
      resp, err := f.get()
      if err != nil {
         return err
      }
      fmt.Println(relay, resp.Status)
      if resp.StatusCode == http.StatusOK {
         break
      }
   }
   return nil
}
