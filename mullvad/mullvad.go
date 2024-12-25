package mullvad

import (
   "os/exec"
   "strings"
   "time"
)

//Connected to ca-tor-wg-103 in Toronto, Canada
//Your connection appears to be from: Canada, Toronto
//
//Connected to ca-tor-wg-103 in Toronto, Canada
//Your connection appears to be from: Canada, Toronto. IPv4: 198.54.132.239

func status(prefix string) error {
   for range time.NewTicker(99 * time.Millisecond).C {
      data, err := exec.Command("mullvad", "status").Output()
      if err != nil {
         return err
      }
      if strings.HasPrefix(string(data), prefix) {
         break
      }
   }
   return nil
}

func connect(location string) error {
   err := exec.Command("mullvad", "relay", "set", "location", location).Run()
   if err != nil {
      return err
   }
   err = exec.Command("mullvad", "connect").Run()
   if err != nil {
      return err
   }
   return status("Connected")
}

func disconnect() error {
   err := exec.Command("mullvad", "disconnect").Run()
   if err != nil {
      return err
   }
   return status("Disconnected")
}
