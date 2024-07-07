package youtube

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
)

func (d DeviceCode) String() string {
   var b strings.Builder
   b.WriteString("1. Go to\n")
   b.WriteString(d.VerificationUrl)
   b.WriteString("\n\n2. Enter this code\n")
   b.WriteString(d.UserCode)
   b.WriteString("\n\n3. Press Enter to continue")
   return b.String()
}

type DeviceCode struct {
   DeviceCode string `json:"device_code"`
   UserCode string `json:"user_code"`
   VerificationUrl string `json:"verification_url"`
}

func (d *DeviceCode) New() error {
   resp, err := http.PostForm(
      "https://oauth2.googleapis.com/device/code",
      url.Values{
         "client_id": {client_id},
         "scope": {"https://www.googleapis.com/auth/youtube"},
      },
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(d)
}
