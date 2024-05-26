package youtube

import (
   "encoding/json"
   "io"
   "net/http"
   "net/url"
)

// YouTube on TV
const (
   client_id = "861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68.apps.googleusercontent.com"
   client_secret = "SboVhoG9s0rNafixCSGGKXAT"
)

func (d DeviceCode) Auth() (*AuthToken, error) {
   res, err := http.PostForm(
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
   defer res.Body.Close()
   var auth AuthToken
   auth.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &auth, nil
}

type AuthToken struct {
   Data []byte
   V struct {
      AccessToken string `json:"access_token"`
      RefreshToken string `json:"refresh_token"`
   }
}

func (a *AuthToken) Unmarshal() error {
   return json.Unmarshal(a.Data, &a.V)
}

func (a *AuthToken) Refresh() error {
   res, err := http.PostForm(
      "https://oauth2.googleapis.com/token",
      url.Values{
         "client_id": {client_id},
         "client_secret": {client_secret},
         "grant_type": {"refresh_token"},
         "refresh_token": {a.V.RefreshToken},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(&a.V)
}
