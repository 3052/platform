package bandcamp

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strconv"
   "time"
)

func (b *BandDetails) New(id int64) error {
   address := func() string {
      b := []byte("http://bandcamp.com/api/mobile/24/band_details?band_id=")
      b = strconv.AppendInt(b, id, 10)
      return string(b)
   }()
   req, err := http.NewRequest("", address, nil)
   if err != nil {
      return err
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(b)
}

type Image struct {
   Crop bool
   Format int
   Height int
   ID int64
   Width int
}

var Images = []Image{
   {ID:0, Width:1500, Height:1500, Format:JPEG},
   {ID:1, Width:1500, Height:1500, Format:PNG},
   {ID:2, Width:350, Height:350, Format:JPEG},
   {ID:3, Width:100, Height:100, Format:JPEG},
   {ID:4, Width:300, Height:300, Format:JPEG},
   {ID:5, Width:700, Height:700, Format:JPEG},
   {ID:6, Width:100, Height:100, Format:JPEG},
   {ID:7, Width:150, Height:150, Format:JPEG},
   {ID:8, Width:124, Height:124, Format:JPEG},
   {ID:9, Width:210, Height:210, Format:JPEG},
   {ID:10, Width:1200, Height:1200, Format:JPEG},
   {ID:11, Width:172, Height:172, Format:JPEG},
   {ID:12, Width:138, Height:138, Format:JPEG},
   {ID:13, Width:380, Height:380, Format:JPEG},
   {ID:14, Width:368, Height:368, Format:JPEG},
   {ID:15, Width:135, Height:135, Format:JPEG},
   {ID:16, Width:700, Height:700, Format:JPEG},
   {ID:20, Width:1024, Height:1024, Format:JPEG},
   {ID:21, Width:120, Height:120, Format:JPEG},
   {ID:22, Width:25, Height:25, Format:JPEG},
   {ID:23, Width:300, Height:300, Format:JPEG},
   {ID:24, Width:300, Height:300, Format:JPEG},
   {ID:25, Width:700, Height:700, Format:JPEG},
   {ID:26, Width:800, Height:600, Format:JPEG, Crop:true},
   {ID:27, Width:715, Height:402, Format:JPEG, Crop:true},
   {ID:28, Width:768, Height:432, Format:JPEG, Crop:true},
   {ID:29, Width:100, Height:75, Format:JPEG, Crop:true},
   {ID:31, Width:1024, Height:1024, Format:PNG},
   {ID:32, Width:380, Height:285, Format:JPEG, Crop:true},
   {ID:33, Width:368, Height:276, Format:JPEG, Crop:true},
   {ID:36, Width:400, Height:300, Format:JPEG, Crop:true},
   {ID:37, Width:168, Height:126, Format:JPEG, Crop:true},
   {ID:38, Width:144, Height:108, Format:JPEG, Crop:true},
   {ID:41, Width:210, Height:210, Format:JPEG},
   {ID:42, Width:50, Height:50, Format:JPEG},
   {ID:43, Width:100, Height:100, Format:JPEG},
   {ID:44, Width:200, Height:200, Format:JPEG},
   {ID:50, Width:140, Height:140, Format:JPEG},
   {ID:65, Width:700, Height:700, Format:JPEG},
   {ID:66, Width:1200, Height:1200, Format:JPEG},
   {ID:67, Width:350, Height:350, Format:JPEG},
   {ID:68, Width:210, Height:210, Format:JPEG},
   {ID:69, Width:700, Height:700, Format:JPEG},
}

// Extension is optional.
func (i Image) URL(art_id int64) string {
   var b []byte
   b = append(b, "http://f4.bcbits.com/img/a"...)
   b = strconv.AppendInt(b, art_id, 10)
   b = append(b, '_')
   b = strconv.AppendInt(b, i.ID, 10)
   return string(b)
}

func (i Item) Tralbum() (*Tralbum, error) {
   switch i.ItemType {
   case "album":
      return new_tralbum('a', i.ItemId)
   case "track":
      return new_tralbum('t', i.ItemId)
   }
   return nil, invalid_type{i.ItemType}
}

func new_tralbum(typ byte, id int) (*Tralbum, error) {
   req, err := http.NewRequest(
      "GET", "http://bandcamp.com/api/mobile/24/tralbum_details", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "band_id": {"1"},
      "tralbum_id": {strconv.Itoa(id)},
      "tralbum_type": {string(typ)},
   }.Encode()
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   tralb := new(Tralbum)
   if err := json.NewDecoder(res.Body).Decode(tralb); err != nil {
      return nil, err
   }
   return tralb, nil
}

type invalid_type struct {
   value string
}

func (i invalid_type) Error() string {
   var b []byte
   b = append(b, "invalid type "...)
   b = strconv.AppendQuote(b, i.value)
   return string(b)
}

const (
   JPEG = iota
   PNG
)

type BandDetails struct {
   Name string
   Discography []Item
}
type AlbumTrack struct {
   TrackNum int64 `json:"track_num"`
   Title string
   BandName string `json:"band_name"`
   StreamingUrl *struct {
      MP3_128 string `json:"mp3-128"`
   } `json:"streaming_url"`
}

func (t Tralbum) Date() time.Time {
   return time.Unix(t.ReleaseDate, 0)
}

type Tralbum struct {
   ArtId int64 `json:"art_id"`
   Title string
   Tracks []AlbumTrack
   ReleaseDate int64 `json:"release_date"`
   TralbumArtist string `json:"tralbum_artist"`
}

type Item struct {
   BandId int64 `json:"band_id"`
   ItemId int `json:"item_id"`
   ItemType string `json:"item_type"`
}

func (i Item) Band() (*BandDetails, error) {
   var band BandDetails
   err := band.New(i.BandId)
   if err != nil {
      return nil, err
   }
   return &band, nil
}
