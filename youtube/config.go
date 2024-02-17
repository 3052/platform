package youtube

import (
   "154.pages.dev/encoding/json"
   "io"
   "net/http"
)

type config struct {
   Innertube_Client_Name string
   Innertube_Client_Version string
}

func (c *config) get() error {
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
