package main

import (
   "41.neocities.org/platform/webDriver"
   "fmt"
   "os"
   "os/exec"
   "time"
)

func after(session *webDriver.SessionState) error {
   after := time.After(9 * time.Second)
   for {
      var err error
      select {
      case <-time.After(time.Second):
         err = session.New()
         if err == nil {
            return nil
         }
      case <-after:
         return err
      }
   }
}

func main() {
   cmd := exec.Command("./geckodriver")
   cmd.Stderr = os.Stderr
   cmd.Stdout = os.Stdout
   err := cmd.Start()
   if err != nil {
      panic(err)
   }
   defer cmd.Wait()
   var session webDriver.SessionState
   err = after(&session)
   if err != nil {
      panic(err)
   }
   err = session.Navigate("http://httpbingo.org/cookies/set?k1=v1&k2=v2")
   if err != nil {
      panic(err)
   }
   cookie, err := session.Cookie()
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", cookie)
}
