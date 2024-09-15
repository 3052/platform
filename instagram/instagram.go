package instagram

import (
   "bytes"
   "encoding/json"
   "net/http"
   "path"
)

func (a *address) Set(text string) error {
   a.shortcode = path.Base(text)
   return nil
}

func (a *address) String() string {
   return a.shortcode
}

type address struct {
   shortcode string
}

const doc_id = 25531498899829322

type media_data struct {
   Data struct {
      XdtShortcodeMedia struct {
         DisplayUrl string `json:"display_url"`
      } `json:"xdt_shortcode_media"`
   }
}

func (m *media_data) New(shortcode string) error {
   var value struct {
      DocId int `json:"doc_id,string"`
      Variables struct {
         Shortcode string `json:"shortcode"`
      } `json:"variables"`
   }
   value.DocId = doc_id
   value.Variables.Shortcode = shortcode
   data, err := json.Marshal(value)
   if err != nil {
      return err
   }
   resp, err := http.Post(
      "https://www.instagram.com/graphql/query",
      "application/json", bytes.NewReader(data),
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(m)
}
