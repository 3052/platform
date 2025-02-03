package soundcloud

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "time"
)

func (t *Transcoding) Media() (*ClientMedia, error) {
   req, err := http.NewRequest("", t.Url, nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "client_id=" + url.QueryEscape(client_id)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   media := &ClientMedia{}
   err = json.NewDecoder(resp.Body).Decode(media)
   if err != nil {
      return nil, err
   }
   return media, nil
}

func (c *ClientTrack) New(id int64) error {
   req, err := http.NewRequest("", "https://api-v2.soundcloud.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = func() string {
      b := []byte("/tracks/")
      b = strconv.AppendInt(b, id, 10)
      return string(b)
   }()
   req.URL.RawQuery = url.Values{
      "client_id": {client_id},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   return json.NewDecoder(resp.Body).Decode(c)
}

const client_id = "KKzJxmw11tYpCs6T24P4uUYhqmjalG6M"

type ClientMedia struct {
   Url string // cf-media.sndcdn.com/QaV7QR1lxpc6.128.mp3
}

// i1.sndcdn.com/artworks-000308141235-7ep8lo-large.jpg
func (c *ClientTrack) Artwork() string {
   if c.ArtworkUrl != "" {
      return c.ArtworkUrl
   }
   return c.User.AvatarUrl
}

func (c *ClientTrack) Resolve(address string) error {
   req, err := http.NewRequest("", "https://api-v2.soundcloud.com/resolve", nil)
   if err != nil {
      return err
   }
   req.URL.RawQuery = url.Values{
      "client_id": {client_id},
      "url": {address},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(c)
}

type ClientTrack struct {
   ArtworkUrl string `json:"artwork_url"`
   DisplayDate time.Time `json:"display_date"`
   Id int64
   Media struct {
      Transcodings []Transcoding
   }
   Title string
   User struct {
      AvatarUrl string `json:"avatar_url"`
      Username string
   }
}

type Transcoding struct {
   Format struct {
      Protocol string
   }
   Url string
}

type hero_artwork struct {
   crop bool
   size string
}

var artworks = []hero_artwork{
   {size: "large"}, // 100x100
   {size: "t120x120"},
   {size: "t1240x260", crop: true},
   {size: "t200x200"},
   {size: "t20x20"},
   {size: "t240x240"},
   {size: "t2480x520", crop: true},
   {size: "t250x250"},
   {size: "t300x300"},
   {size: "t40x40"},
   {size: "t47x47"},
   {size: "t500x"},
   {size: "t500x500"},
   {size: "t50x50"},
   {size: "t60x60"},
   {size: "t67x67"},
   {size: "t80x80"},
   {size: "tx250"},
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
   text, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   sep := []byte(` id="__NEXT_DATA__" type="application/json">`)
   _, text, _ = bytes.Cut(text, sep)
   return json.NewDecoder(bytes.NewReader(text)).Decode(n)
}
