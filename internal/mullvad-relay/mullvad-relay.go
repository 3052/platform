package main

import (
   "41.neocities.org/platform/mullvad"
   "fmt"
   "io"
   "net/http"
)

func (f *flags) connect(servers mullvad.ServerList) error {
   for server := range servers.Seq(f.location) {
      fmt.Print(server)
      err := mullvad.Connect(server)
      switch err {
      case nil:
         fmt.Print(" GET")
         resp, err := f.get()
         if err != nil {
            return err
         }
         fmt.Printf(" %v\n", resp.Status)
         if resp.StatusCode == http.StatusOK {
            return nil
         }
      case mullvad.ErrTimeout:
         fmt.Printf(" %v\n", err)
      default:
         return err
      }
   }
   return io.EOF
}

func (f *flags) relay() error {
   defer mullvad.Disconnect()
   var servers mullvad.ServerList
   err := servers.OpenVpn()
   if err != nil {
      return err
   }
   err = f.connect(servers)
   switch err {
   case nil:
      return nil
   case io.EOF:
      var servers mullvad.ServerList
      err = servers.WireGuard()
      if err != nil {
         return err
      }
      return f.connect(servers)
   default:
      return err
   }
}

func (f *flags) get() (*http.Response, error) {
   resp, err := http.Get(f.address)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return resp, nil
}
