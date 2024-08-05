package youtube

import (
   "encoding/json"
   "net/http"
   "net/url"
)

// YouTube on TV
const (
   client_id = "861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68.apps.googleusercontent.com"
   client_secret = "SboVhoG9s0rNafixCSGGKXAT"
)

func (a *AuthToken) Refresh() error {
   resp, err := http.PostForm(
      "https://oauth2.googleapis.com/token", url.Values{
         "client_id": {client_id},
         "client_secret": {client_secret},
         "grant_type": {"refresh_token"},
         "refresh_token": {a.RefreshToken},
      },
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(a)
}

func (a *AuthToken) Marshal() ([]byte, error) {
   return json.Marshal(a)
}

type AuthToken struct {
   AccessToken string `json:"access_token"`
   RefreshToken string `json:"refresh_token"`
}

func (a *AuthToken) Unmarshal(text []byte) error {
   return json.Unmarshal(text, a)
}

func (d *DeviceCode) Token() (*AuthToken, error) {
   resp, err := http.PostForm(
      "https://oauth2.googleapis.com/token", url.Values{
         "client_id": {client_id},
         "client_secret": {client_secret},
         "device_code": {d.DeviceCode},
         "grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
      },
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   token := &AuthToken{}
   err = json.NewDecoder(resp.Body).Decode(token)
   if err != nil {
      return nil, err
   }
   return token, nil
}
