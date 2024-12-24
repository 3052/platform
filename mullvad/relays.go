package mullvad

import "net/http"

func relays() (*http.Response, error) {
   return http.Get("https://api.mullvad.net/public/relays/v1")
}
