package soundcloud

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
)

func (n *next_data) New() error {
   resp, err := http.Get("https://m.soundcloud.com")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   text, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   sep := []byte(` id="__NEXT_DATA__" type="application/json">`)
   _, text, _ = bytes.Cut(text, sep)
   return json.NewDecoder(bytes.NewReader(text)).Decode(n)
}

type next_data struct {
   RuntimeConfig struct {
      ClientId string `json:"clientId"`
   } `json:"runtimeConfig"`
}
