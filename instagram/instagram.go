package instagram

import (
   "bytes"
   "encoding/json"
   "net/http"
   "net/url"
   "path"
)

func (s *ShortcodeMedia) DisplayUrls() []Url {
   if v := s.EdgeSidecarToChildren; v != nil {
      var media []Url
      for _, medium := range v.Edges {
         media = append(media, medium.Node.DisplayUrl)
      }
      return media
   }
   return []Url{s.DisplayUrl}
}

type Url struct {
   Url url.URL
}

func (b *Url) UnmarshalText(data []byte) error {
   return b.Url.UnmarshalBinary(data)
}

const doc_id = 25531498899829322

type ShortcodeMedia struct {
   DisplayUrl Url `json:"display_url"`
   EdgeSidecarToChildren *struct {
      Edges []struct {
         Node struct {
            DisplayUrl Url `json:"display_url"`
         }
      }
   } `json:"edge_sidecar_to_children"`
}

func (s ShortCode) String() string {
   return s()
}

func (s *ShortCode) Set(data string) error {
   *s = func() string {
      return path.Base(data)
   }
   return nil
}

type ShortCode func() string

func (s ShortCode) Media() (*ShortcodeMedia, error) {
   var value struct {
      DocId int `json:"doc_id,string"`
      Variables struct {
         Shortcode string `json:"shortcode"`
      } `json:"variables"`
   }
   value.DocId = doc_id
   value.Variables.Shortcode = s()
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
