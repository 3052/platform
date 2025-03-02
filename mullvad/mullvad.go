package mullvad

import (
   "log"
   "net/http"
   "os/exec"
)

func command(name string, arg ...string) *exec.Cmd {
   cmd := exec.Command(name, arg...)
   log.Print(cmd.Args)
   return cmd
}

type Transport http.Transport

func Connect() error {
   return command("mullvad", "connect", "-w").Run()
}

func Disconnect() error {
   return command("mullvad", "disconnect", "-w").Run()
}

var connected bool

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   var err error
   if req.Header.Get("vpn") != "" {
      if !connected {
         err = Connect()
         connected = true
      }
   } else if connected {
      err = Disconnect()
      connected = false
   }
   if err != nil {
      return nil, err
   }
   if req.Header.Get("silent") == "" {
      log.Println(req.Method, req.URL)
   }
   return (*http.Transport)(t).RoundTrip(req)
}
