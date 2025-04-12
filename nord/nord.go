package nord

import "net/http"

func servers() (*http.Response, error) {
   return http.Get("https://api.nordvpn.com/v1/servers")
}
