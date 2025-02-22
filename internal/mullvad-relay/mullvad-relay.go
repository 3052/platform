package main

import (
   "41.neocities.org/platform/mullvad"
   "flag"
   "fmt"
   "io"
   "net/http"
)

func (f *flags) connect(relays mullvad.PublicRelays) error {
   for relay := range relays.Hostname(f.location) {
      fmt.Print(relay)
      err := mullvad.Location(relay).Connect()
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

type flags struct {
   address  string
   location string
}

func main() {
   var f flags
   flag.StringVar(&f.address, "a", "", "address")
   flag.StringVar(&f.location, "n", "", "location")
   flag.Parse()
   if f.address != "" {
      err := f.relay()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

func (f *flags) get() (*http.Response, error) {
   resp, err := http.Get(f.address)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   _, err = io.Copy(io.Discard, resp.Body)
   if err != nil {
      return nil, err
   }
   return resp, nil
}

func (f *flags) relay() error {
   defer mullvad.Disconnect()
   var relays mullvad.PublicRelays
   err := relays.OpenVpn()
   if err != nil {
      return err
   }
   err = f.connect(relays)
   switch err {
   case nil:
      return nil
   case io.EOF:
      var relays mullvad.PublicRelays
      err = relays.WireGuard()
      if err != nil {
         return err
      }
      return f.connect(relays)
   default:
      return err
   }
}
