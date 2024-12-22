package bandcamp

import (
   "41.neocities.org/text"
   "encoding/json"
   "encoding/xml"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "time"
)

func (b *BandDetails) New(id int) error {
   req, err := http.NewRequest("", "http://bandcamp.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/api/mobile/24/band_details")
   req.URL.RawQuery = "band_id=" + strconv.Itoa(id)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(b)
}

func (r *ReportParams) New(address string) error {
   resp, err := http.Get(address)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   _, data, _ = text.CutBefore(data, []byte(`<p id="report-account-vm"`))
   var p struct {
      DataTouReportParams []byte `xml:"data-tou-report-params,attr"`
   }
   err = xml.Unmarshal(data, &p)
   if err != nil {
      return err
   }
   return json.Unmarshal(p.DataTouReportParams, r)
}

func (r *ReportParams) Band() (*BandDetails, error) {
   var band BandDetails
   err := band.New(r.Aid)
   if err != nil {
      return nil, err
   }
   return &band, nil
}

type ReportParams struct {
   Aid int64 `json:"a_id"`
   Iid int `json:"i_id"`
   Itype string `json:"i_type"`
}

const (
   Jpeg = iota
   Png
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

func (t *Tralbum) Date() time.Time {
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

func (i *Item) Band() (*BandDetails, error) {
   var band BandDetails
   err := band.New(i.BandId)
   if err != nil {
      return nil, err
   }
   return &band, nil
}

type Image struct {
   Crop bool
   Format int
   Height int
   ID int64
   Width int
}

var Images = []Image{
   {ID:0, Width:1500, Height:1500, Format:Jpeg},
   {ID:1, Width:1500, Height:1500, Format:Png},
   {ID:2, Width:350, Height:350, Format:Jpeg},
   {ID:3, Width:100, Height:100, Format:Jpeg},
   {ID:4, Width:300, Height:300, Format:Jpeg},
   {ID:5, Width:700, Height:700, Format:Jpeg},
   {ID:6, Width:100, Height:100, Format:Jpeg},
   {ID:7, Width:150, Height:150, Format:Jpeg},
   {ID:8, Width:124, Height:124, Format:Jpeg},
   {ID:9, Width:210, Height:210, Format:Jpeg},
   {ID:10, Width:1200, Height:1200, Format:Jpeg},
   {ID:11, Width:172, Height:172, Format:Jpeg},
   {ID:12, Width:138, Height:138, Format:Jpeg},
   {ID:13, Width:380, Height:380, Format:Jpeg},
   {ID:14, Width:368, Height:368, Format:Jpeg},
   {ID:15, Width:135, Height:135, Format:Jpeg},
   {ID:16, Width:700, Height:700, Format:Jpeg},
   {ID:20, Width:1024, Height:1024, Format:Jpeg},
   {ID:21, Width:120, Height:120, Format:Jpeg},
   {ID:22, Width:25, Height:25, Format:Jpeg},
   {ID:23, Width:300, Height:300, Format:Jpeg},
   {ID:24, Width:300, Height:300, Format:Jpeg},
   {ID:25, Width:700, Height:700, Format:Jpeg},
   {ID:26, Width:800, Height:600, Format:Jpeg, Crop:true},
   {ID:27, Width:715, Height:402, Format:Jpeg, Crop:true},
   {ID:28, Width:768, Height:432, Format:Jpeg, Crop:true},
   {ID:29, Width:100, Height:75, Format:Jpeg, Crop:true},
   {ID:31, Width:1024, Height:1024, Format:Png},
   {ID:32, Width:380, Height:285, Format:Jpeg, Crop:true},
   {ID:33, Width:368, Height:276, Format:Jpeg, Crop:true},
   {ID:36, Width:400, Height:300, Format:Jpeg, Crop:true},
   {ID:37, Width:168, Height:126, Format:Jpeg, Crop:true},
   {ID:38, Width:144, Height:108, Format:Jpeg, Crop:true},
   {ID:41, Width:210, Height:210, Format:Jpeg},
   {ID:42, Width:50, Height:50, Format:Jpeg},
   {ID:43, Width:100, Height:100, Format:Jpeg},
   {ID:44, Width:200, Height:200, Format:Jpeg},
   {ID:50, Width:140, Height:140, Format:Jpeg},
   {ID:65, Width:700, Height:700, Format:Jpeg},
   {ID:66, Width:1200, Height:1200, Format:Jpeg},
   {ID:67, Width:350, Height:350, Format:Jpeg},
   {ID:68, Width:210, Height:210, Format:Jpeg},
   {ID:69, Width:700, Height:700, Format:Jpeg},
}

// Extension is optional.
func (i *Image) URL(art_id int64) string {
   var b []byte
   b = append(b, "http://f4.bcbits.com/img/a"...)
   b = strconv.AppendInt(b, art_id, 10)
   b = append(b, '_')
   b = strconv.AppendInt(b, i.ID, 10)
   return string(b)
}

func new_tralbum(typ byte, id int) (*Tralbum, error) {
   req, err := http.NewRequest(
      "", "http://bandcamp.com/api/mobile/24/tralbum_details", nil,
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "band_id": {"1"},
      "tralbum_id": {strconv.Itoa(id)},
      "tralbum_type": {string(typ)},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   album := &Tralbum{}
   if err := json.NewDecoder(resp.Body).Decode(album); err != nil {
      return nil, err
   }
   return album, nil
}

type invalid_type struct {
   data string
}

func (r *ReportParams) Tralbum() (*Tralbum, error) {
   switch r.Itype {
   case "a":
      return new_tralbum('a', r.Iid)
   case "t":
      return new_tralbum('t', r.Iid)
   }
   return nil, invalid_type{r.Itype}
}

func (i *Item) Tralbum() (*Tralbum, error) {
   switch i.ItemType {
   case "album":
      return new_tralbum('a', i.ItemId)
   case "track":
      return new_tralbum('t', i.ItemId)
   }
   return nil, invalid_type{i.ItemType}
}

func (i invalid_type) Error() string {
   var data []byte
   data = append(data, "invalid type "...)
   data = strconv.AppendQuote(data, i.data)
   return string(data)
}
