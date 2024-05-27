package bandcamp

import (
   "154.pages.dev/text"
   "encoding/json"
   "encoding/xml"
   "io"
   "net/http"
)

func (r *ReportParams) New(address string) error {
   res, err := http.Get(address)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   data, err := io.ReadAll(res.Body)
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

func (p ReportParams) Band() (*BandDetails, error) {
   var band BandDetails
   err := band.New(p.Aid)
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

func (p ReportParams) Tralbum() (*Tralbum, error) {
   switch p.Itype {
   case "a":
      return new_tralbum('a', p.Iid)
   case "t":
      return new_tralbum('t', p.Iid)
   }
   return nil, invalid_type{p.Itype}
}
