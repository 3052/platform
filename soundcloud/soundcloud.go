package soundcloud

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "net/url"
   "time"
)

const client_id = "KKzJxmw11tYpCs6T24P4uUYhqmjalG6M"

type Resolve struct {
   ArtworkUrl string `json:"artwork_url"`
   DisplayDate time.Time `json:"display_date"`
   Id int64
   Title string
   User struct {
      AvatarUrl string `json:"avatar_url"`
      Username string
   }
}

func (r *Resolve) New(url1 string) error {
   req, _ := http.NewRequest("", "https://api-v2.soundcloud.com/resolve", nil)
   req.URL.RawQuery = url.Values{
      "client_id": {client_id},
      "url": {url1},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(r)
}

// i1.sndcdn.com/artworks-000308141235-7ep8lo-large.jpg
func (r *Resolve) Artwork() string {
   if r.ArtworkUrl != "" {
      return r.ArtworkUrl
   }
   return r.User.AvatarUrl
}

type next_data struct {
   RuntimeConfig struct {
      ClientId string `json:"clientId"`
   } `json:"runtimeConfig"`
}

func (n *next_data) New() error {
   resp, err := http.Get("https://m.soundcloud.com")
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   sep := []byte(` id="__NEXT_DATA__" type="application/json">`)
   _, data, _ = bytes.Cut(data, sep)
   return json.NewDecoder(bytes.NewReader(data)).Decode(n)
}
