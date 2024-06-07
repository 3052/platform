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
   res, err := http.Get(address)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(c)
}

type ClientTrack struct {
   ArtworkUrl string `json:"artwork_url"`
   DisplayDate string `json:"display_date"`
   ID int64
   Media struct {
      Transcodings []struct {
         Format struct {
            Protocol string
         }
         URL string
      }
   }
   Title string
   User struct {
      AvatarUrl string `json:"avatar_url"`
      Username string
   }
}

// Also available is "hls", but all transcodings are quality "sq".
// Same for "api-mobile.soundcloud.com".
func (t ClientTrack) Progressive() (*Media, error) {
   ref := func() string {
      for _, code := range t.Media.Transcodings {
         if code.Format.Protocol == "progressive" {
            return code.URL
         }
      }
      return ""
   }
   req, err := http.NewRequest("", ref(), nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "client_id=" + client_id
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   med := new(Media)
   if err := json.NewDecoder(res.Body).Decode(med); err != nil {
      return nil, err
   }
   return med, nil
}

type Media struct {
   URL string // cf-media.sndcdn.com/QaV7QR1lxpc6.128.mp3
}

type Image struct {
   Size string
   Crop bool
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

// i1.sndcdn.com/artworks-000308141235-7ep8lo-large.jpg
func (t ClientTrack) Artwork() string {
   if t.ArtworkUrl == "" {
      t.ArtworkUrl = t.User.AvatarUrl
   }
   return strings.Replace(t.ArtworkUrl, "large", "t500x500", 1)
}

func (t ClientTrack) Time() (time.Time, error) {
   return time.Parse(time.RFC3339, t.DisplayDate)
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
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var s struct {
      Track ClientTrack
   }
   err = json.NewDecoder(res.Body).Decode(&s)
   if err != nil {
      return nil, err
   }
   return &s.Track, nil
}
