package soundcloud

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

const client_id = "iZIs9mchVcX5lhVRyQGGAYlNPVldzAoX"

type next_data struct {
   RuntimeConfig struct {
      ClientId string `json:"clientId"`
   } `json:"runtimeConfig"`
}

type hero_artwork struct {
   crop bool
   size string
}

var Artwork = []hero_artwork{
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

type ClientTrack struct {
   ArtworkUrl string `json:"artwork_url"`
   DisplayDate string `json:"display_date"`
   Id int64
   Media struct {
      Transcodings []struct {
         Format struct {
            Protocol string
         }
         Url string
      }
   }
   Title string
   User struct {
      AvatarUrl string `json:"avatar_url"`
      Username string
   }
}

type Media struct {
   Url string // cf-media.sndcdn.com/QaV7QR1lxpc6.128.mp3
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

func (c *ClientTrack) New(id int64) error {
   address := func() string {
      b := []byte("https://api-v2.soundcloud.com/tracks/")
      b = strconv.AppendInt(b, id, 10)
      b = append(b, "?client_id="...)
      b = append(b, client_id...)
      return string(b)
   }()
   resp, err := http.Get(address)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(c)
}

func Resolve(address string) (*ClientTrack, error) {
   req, err := http.NewRequest("", "https://api-v2.soundcloud.com/resolve", nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "client_id": {client_id},
      "url": {address},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var resolve struct {
      Track ClientTrack
   }
   err = json.NewDecoder(resp.Body).Decode(&resolve)
   if err != nil {
      return nil, err
   }
   return &resolve.Track, nil
}

// Also available is "hls", but all transcodings are quality "sq".
// Same for "api-mobile.soundcloud.com".
func (c ClientTrack) Progressive() (*Media, error) {
   address := func() string {
      for _, code := range c.Media.Transcodings {
         if code.Format.Protocol == "progressive" {
            return code.Url
         }
      }
      return ""
   }
   req, err := http.NewRequest("", address(), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "client_id=" + client_id
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   data := &Media{}
   if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
      return nil, err
   }
   return data, nil
}

// i1.sndcdn.com/artworks-000308141235-7ep8lo-large.jpg
func (c ClientTrack) Artwork() string {
   if c.ArtworkUrl == "" {
      c.ArtworkUrl = c.User.AvatarUrl
   }
   return strings.Replace(c.ArtworkUrl, "large", "t500x500", 1)
}

func (c ClientTrack) Time() (time.Time, error) {
   return time.Parse(time.RFC3339, c.DisplayDate)
}
