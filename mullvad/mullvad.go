package mullvad

import (
   "os/exec"
   "strings"
   "time"
)

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

func status(prefix string) error {
   for range time.Tick(99 * time.Millisecond) {
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
