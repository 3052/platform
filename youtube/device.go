package youtube

import (
   "encoding/json"
   "net/http"
   "net/url"
)

func (d *DeviceCode) Post() error {
   res, err := http.PostForm(
      "https://oauth2.googleapis.com/device/code",
      url.Values{
         "client_id": {client_id},
         "scope": {"https://www.googleapis.com/auth/youtube"},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(d)
}
