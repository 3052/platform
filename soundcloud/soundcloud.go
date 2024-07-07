package soundcloud

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

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

//////////

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
   med := new(Media)
   if err := json.NewDecoder(resp.Body).Decode(med); err != nil {
      return nil, err
   }
   return med, nil
}

type Image struct {
   Crop bool
   Size string
}

var Images = []Image{
   {Size: "t120x120"},
   {Size: "t1240x260", Crop: true},
   {Size: "t200x200"},
   {Size: "t20x20"},
   {Size: "t240x240"},
   {Size: "t2480x520", Crop: true},
   {Size: "t250x250"},
   {Size: "t300x300"},
   {Size: "t40x40"},
   {Size: "t47x47"},
   {Size: "t500x"},
   {Size: "t500x500"},
   {Size: "t50x50"},
   {Size: "t60x60"},
   {Size: "t67x67"},
   {Size: "t80x80"},
   {Size: "tx250"},
}

const client_id = "iZIs9mchVcX5lhVRyQGGAYlNPVldzAoX"

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
