package instagram

import (
   "bytes"
   "encoding/json"
   "net/http"
   "path"
)

const doc_id = 25531498899829322

func (a *Address) Set(text string) error {
   a.Shortcode = path.Base(text)
   return nil
}

func (a *Address) String() string {
   return a.Shortcode
}

type Address struct {
   Shortcode string
}

func (a *Address) Media() (*ShortcodeMedia, error) {
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
   var resp_value struct {
      Data struct {
         XdtShortcodeMedia ShortcodeMedia `json:"xdt_shortcode_media"`
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&resp_value)
   if err != nil {
      return nil, err
   }
   return &resp_value.Data.XdtShortcodeMedia, nil
}

type ShortcodeMedia struct {
   DisplayUrl string `json:"display_url"`
   EdgeSidecarToChildren *struct {
      Edges []struct {
         Node struct {
            DisplayUrl string `json:"display_url"`
         }
      }
   } `json:"edge_sidecar_to_children"`
}

func (s *ShortcodeMedia) DisplayUrls() []string {
   if v := s.EdgeSidecarToChildren; v != nil {
      var media []string
      for _, medium := range v.Edges {
         media = append(media, medium.Node.DisplayUrl)
      }
      return media
   }
   return []string{s.DisplayUrl}
}
