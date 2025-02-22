package mullvad

import (
   "bytes"
   "errors"
   "io"
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

type Transport struct{}

func (Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   err := Connect()
   if err != nil {
      return nil, err
   }
   defer Disconnect()
   log.Println(req.Method, req.URL)
   resp, err := http.DefaultTransport.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   resp.Body = io.NopCloser(bytes.NewReader(data))
   return resp, nil
}
