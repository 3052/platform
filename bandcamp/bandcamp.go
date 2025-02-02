package bandcamp

import (
   "bytes"
   "encoding/json"
   "encoding/xml"
   "io"
   "net/http"
   "net/url"
   "strconv"
   "time"
)

func (r *report_params) New(url0 string) error {
   resp, err := http.Get(url0)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   _, data, _ = cut_before(data, []byte(`<p id="report-account-vm"`))
   var p struct {
      DataTouReportParams []byte `xml:"data-tou-report-params,attr"`
   }
   err = xml.Unmarshal(data, &p)
   if err != nil {
      return err
   }
   return json.Unmarshal(p.DataTouReportParams, r)
}

func (r *report_params) tralbum() (*tralbum, bool) {
   switch r.Itype {
   case "a":
      return &tralbum{r.Iid, 'a'}, true
   case "t":
      return &tralbum{r.Iid, 't'}, true
   }
   return nil, false
}

func (b *band_details) New(band_id int64) error {
   req, _ := http.NewRequest("", "http://bandcamp.com", nil)
   req.URL.Path = "/api/mobile/24/band_details"
   req.URL.RawQuery = "band_id=" + strconv.FormatInt(band_id, 10)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(b)
}

type band_details struct {
   Discography []recording
   Name        string
}

type report_params struct {
   Aid   int64  `json:"a_id"`
   Iid   int    `json:"i_id"`
   Itype string `json:"i_type"`
}

type recording struct {
   BandId   int64  `json:"band_id"`
   ItemId   int    `json:"item_id"`
   ItemType string `json:"item_type"`
}

type tralbum_details struct {
   ArtId         int64 `json:"art_id"`
   ReleaseDate   int64  `json:"release_date"`
   Title         string
   Tracks        []track
   TralbumArtist string `json:"tralbum_artist"`
}

func cut_before(s, sep []byte) ([]byte, []byte, bool) {
   i := bytes.Index(s, sep)
   if i >= 0 {
      return s[:i], s[i:], true
   }
   return s, nil, false
}

const (
   jpeg = iota
   png
)

func (r *recording) tralbum() (*tralbum, bool) {
   switch r.ItemType {
   case "album":
      return &tralbum{r.ItemId, 'a'}, true
   case "track":
      return &tralbum{r.ItemId, 't'}, true
   }
   return nil, false
}

func (r *report_params) Band() (*band_details, error) {
   var band band_details
   err := band.New(r.Aid)
   if err != nil {
      return nil, err
   }
   return &band, nil
}

func (r *recording) Band() (*band_details, error) {
   var band band_details
   err := band.New(r.BandId)
   if err != nil {
      return nil, err
   }
   return &band, nil
}

type tralbum struct {
   tralbum_id int
   tralbum_type byte
}

func (t *tralbum) tralbum() (*tralbum_details, error) {
   req, _ := http.NewRequest("", "http://bandcamp.com", nil)
   req.URL.Path = "/api/mobile/24/tralbum_details"
   req.URL.RawQuery = url.Values{
      "band_id":      {"1"},
      "tralbum_id":   {strconv.Itoa(t.tralbum_id)},
      "tralbum_type": {string(t.tralbum_type)},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   detail := &tralbum_details{}
   if err := json.NewDecoder(resp.Body).Decode(detail); err != nil {
      return nil, err
   }
   return detail, nil
}

func (t *tralbum_details) Time() time.Time {
   return time.Unix(t.ReleaseDate, 0)
}

type track struct {
   BandName     string `json:"band_name"`
   StreamingUrl *struct {
      Mp3_128 string `json:"mp3-128"`
   } `json:"streaming_url"`
   Title        string
   TrackNum     int64 `json:"track_num"`
}

// Extension is optional.
func (i *image) Url(art_id int64) string {
   var b []byte
   b = append(b, "http://f4.bcbits.com/img/a"...)
   b = strconv.AppendInt(b, art_id, 10)
   b = append(b, '_')
   b = strconv.AppendInt(b, i.Id, 10)
   return string(b)
}

type image struct {
   Crop   bool
   Format int
   Height int
   Id     int64
   Width  int
}

var images = []image{
   {Id: 0, Width: 1500, Height: 1500, Format: jpeg},
   {Id: 1, Width: 1500, Height: 1500, Format: png},
   {Id: 2, Width: 350, Height: 350, Format: jpeg},
   {Id: 3, Width: 100, Height: 100, Format: jpeg},
   {Id: 4, Width: 300, Height: 300, Format: jpeg},
   {Id: 5, Width: 700, Height: 700, Format: jpeg},
   {Id: 6, Width: 100, Height: 100, Format: jpeg},
   {Id: 7, Width: 150, Height: 150, Format: jpeg},
   {Id: 8, Width: 124, Height: 124, Format: jpeg},
   {Id: 9, Width: 210, Height: 210, Format: jpeg},
   {Id: 10, Width: 1200, Height: 1200, Format: jpeg},
   {Id: 11, Width: 172, Height: 172, Format: jpeg},
   {Id: 12, Width: 138, Height: 138, Format: jpeg},
   {Id: 13, Width: 380, Height: 380, Format: jpeg},
   {Id: 14, Width: 368, Height: 368, Format: jpeg},
   {Id: 15, Width: 135, Height: 135, Format: jpeg},
   {Id: 16, Width: 700, Height: 700, Format: jpeg},
   {Id: 20, Width: 1024, Height: 1024, Format: jpeg},
   {Id: 21, Width: 120, Height: 120, Format: jpeg},
   {Id: 22, Width: 25, Height: 25, Format: jpeg},
   {Id: 23, Width: 300, Height: 300, Format: jpeg},
   {Id: 24, Width: 300, Height: 300, Format: jpeg},
   {Id: 25, Width: 700, Height: 700, Format: jpeg},
   {Id: 26, Width: 800, Height: 600, Format: jpeg, Crop: true},
   {Id: 27, Width: 715, Height: 402, Format: jpeg, Crop: true},
   {Id: 28, Width: 768, Height: 432, Format: jpeg, Crop: true},
   {Id: 29, Width: 100, Height: 75, Format: jpeg, Crop: true},
   {Id: 31, Width: 1024, Height: 1024, Format: png},
   {Id: 32, Width: 380, Height: 285, Format: jpeg, Crop: true},
   {Id: 33, Width: 368, Height: 276, Format: jpeg, Crop: true},
   {Id: 36, Width: 400, Height: 300, Format: jpeg, Crop: true},
   {Id: 37, Width: 168, Height: 126, Format: jpeg, Crop: true},
   {Id: 38, Width: 144, Height: 108, Format: jpeg, Crop: true},
   {Id: 41, Width: 210, Height: 210, Format: jpeg},
   {Id: 42, Width: 50, Height: 50, Format: jpeg},
   {Id: 43, Width: 100, Height: 100, Format: jpeg},
   {Id: 44, Width: 200, Height: 200, Format: jpeg},
   {Id: 50, Width: 140, Height: 140, Format: jpeg},
   {Id: 65, Width: 700, Height: 700, Format: jpeg},
   {Id: 66, Width: 1200, Height: 1200, Format: jpeg},
   {Id: 67, Width: 350, Height: 350, Format: jpeg},
   {Id: 68, Width: 210, Height: 210, Format: jpeg},
   {Id: 69, Width: 700, Height: 700, Format: jpeg},
}
