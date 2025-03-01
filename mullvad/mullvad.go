package mullvad

import (
   "errors"
   "log"
   "net/http"
   "os/exec"
   "strings"
   "time"
)

var ErrTimeout = errors.New("timeout")

func status(prefix string) error {
   after := time.After(30 * time.Second)
   for {
      select {
      case <-time.After(time.Second):
         data, err := command("mullvad", "status").Output()
         if err != nil {
            return err
         }
         data1 := string(data)
         if strings.HasPrefix(data1, prefix) {
            if strings.Contains(data1, " IPv4:") {
               return nil
            }
         }
      case <-after:
         return ErrTimeout
      }
   }
}

func Relay(location ...string) error {
   arg := []string{"relay", "set", "location"}
   arg = append(arg, location...)
   data, err := command("mullvad", arg...).CombinedOutput()
   if err != nil {
      return errors.New(string(data))
   }
   return nil
}

func command(name string, arg ...string) *exec.Cmd {
   cmd := exec.Command(name, arg...)
   log.Print(cmd.Args)
   return cmd
}

func Disconnect() error {
   err := command("mullvad", "disconnect").Run()
   if err != nil {
      return err
   }
   return status("Disconnected")
}

func Connect() error {
   err := command("mullvad", "connect").Run()
   if err != nil {
      return err
   }
   return status("Connected")
}

type Transport http.Transport

var connected bool

// - disable http/2
// - enable log
// - enable mullvad
// - enable proxy
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
   log.Println(req.Method, req.URL)
   return (*http.Transport)(t).RoundTrip(req)
}
