package main

import (
   "41.neocities.org/platform/webDriver"
   "log"
   "os"
   "os/exec"
   "time"
)

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
   err = wait_session(&session)
   if err != nil {
      panic(err)
   }
   err = session.Navigate("https://www.sky.ch")
   if err != nil {
      panic(err)
   }
   cookie, err := wait_cookie(&session)
   if err != nil {
      panic(err)
   }
   log.Print(cookie)
}

func wait_cookie(session *webDriver.SessionState) (string, error) {
   for {
      time.Sleep(4 * time.Second)
      log.Print("session.Cookie")
      cookie, err := session.Cookie()
      if err != nil {
         return "", err
      }
      for _, value := range cookie.Value {
         if value.Name == "sky-auth-token" {
            return value.Value, nil
         }
      }
   }
}

func wait_session(session *webDriver.SessionState) error {
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
