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

func cut_before(s, sep []byte) ([]byte, []byte, bool) {
   i := bytes.Index(s, sep)
   if i >= 0 {
      return s[:i], s[i:], true
   }
   return s, nil, false
}

type ReportParams struct {
   Aid   int64  `json:"a_id"`
   Iid   int    `json:"i_id"`
   Itype string `json:"i_type"`
}

func (r *ReportParams) New(url0 string) error {
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

func (r *ReportParams) Tralbum() (*Tralbum, bool) {
   switch r.Itype {
   case "a":
      return &Tralbum{r.Iid, 'a'}, true
   case "t":
      return &Tralbum{r.Iid, 't'}, true
   }
   return nil, false
}

type Tralbum struct {
   Id int
   Type byte
}

func (t *Tralbum) Tralbum() (*TralbumDetails, error) {
   req, _ := http.NewRequest("", "http://bandcamp.com", nil)
   req.URL.Path = "/api/mobile/24/tralbum_details"
   req.URL.RawQuery = url.Values{
      "band_id":      {"1"},
      "tralbum_id":   {strconv.Itoa(t.Id)},
      "tralbum_type": {string(t.Type)},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   detail := &TralbumDetails{}
   if err := json.NewDecoder(resp.Body).Decode(detail); err != nil {
      return nil, err
   }
   return detail, nil
}

func (t *TralbumDetails) Time() time.Time {
   return time.Unix(t.ReleaseDate, 0)
}

type TralbumDetails struct {
   ArtId         int64 `json:"art_id"`
   ReleaseDate   int64  `json:"release_date"`
   Title         string
   TralbumArtist string `json:"tralbum_artist"`
}
