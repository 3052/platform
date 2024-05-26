package youtube

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
)

type innertube struct {
   InnertubeClientName string `json:"innertube_client_name"`
   InnertubeClientVersion string `json:"innertube_client_version"`
}

func (i *innertube) New() error {
   res, err := http.Get("https://www.youtube.com")
   if err != nil {
      return err
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   _, text, _ = bytes.Cut(text, []byte("\nytcfg.set("))
   return json.NewDecoder(bytes.NewReader(text)).Decode(i)
}
