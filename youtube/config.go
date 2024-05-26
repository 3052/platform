package youtube

import (
   "io"
   "net/http"
)

type config struct {
   InnertubeClientName string `json:"innertube_client_name"`
   InnertubeClientVersion string `json:"innertube_client_version"`
}

func (c *config) New() error {
   res, err := http.Get("https://www.youtube.com")
   if err != nil {
      return err
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   _, text = json.Cut(text, []byte("\nytcfg.set("), nil)
   return json.Decode(text, c)
}
