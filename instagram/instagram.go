package instagram

import (
   "bytes"
   "encoding/json"
   "net/http"
   "path"
)

func (m *MediaData) String() string {
   return m.Data.XdtShortcodeMedia.DisplayUrl
}

type MediaData struct {
   Data struct {
      XdtShortcodeMedia struct {
         DisplayUrl string `json:"display_url"`
      } `json:"xdt_shortcode_media"`
   }
}

func (a *Address) Set(text string) error {
   a.Shortcode = path.Base(text)
   return nil
}

func (a *Address) String() string {
   return a.Shortcode
}

const doc_id = 25531498899829322

type Address struct {
   Shortcode string
}

func (a *Address) Media() (*MediaData, error) {
   var value struct {
      DocId int `json:"doc_id,string"`
      Variables struct {
         Shortcode string `json:"shortcode"`
      } `json:"variables"`
   }
   value.DocId = doc_id
   value.Variables.Shortcode = a.Shortcode
   data, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   resp, err := http.Post(
      "https://www.instagram.com/graphql/query",
      "application/json", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   media := &MediaData{}
   err = json.NewDecoder(resp.Body).Decode(media)
   if err != nil {
      return nil, err
   }
   return media, nil
}
