package justwatch

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

// map[string]map[string][]string

func (s LocaleState) String(o Offer) string {
   var b strings.Builder
   b.WriteString("url = ")
   b.WriteString(o.StandardWebUrl)
   b.WriteString("\nmonetization = ")
   b.WriteString(o.MonetizationType)
   b.WriteString("\ncountry = ")
   b.WriteString(s.CountryName)
   return b.String()
}

// `presentationType` data seems to be incorrect in some cases. For example,
// JustWatch reports this as SD: fetchtv.com.au/movie/details/19285
// when the site itself reports as HD
type Offer struct {
   MonetizationType string
   StandardWebUrl string
}

func (o Offer) Stream() bool {
   switch o.MonetizationType {
   case "BUY", "RENT":
      return false
   }
   return true
}

func (t LangTag) Offers(s *LocaleState) ([]Offer, error) {
   body, err := func() ([]byte, error) {
      var v struct {
         Variables struct {
            Country string `json:"country"`
            FullPath string `json:"fullPath"`
         }
         Query string
      }
      v.Query = graphql_compact(title_details)
      v.Variables.FullPath = t.Href
      v.Variables.Country = s.Country
      return json.Marshal(v)
   }()
   if err != nil {
      return nil, err
   }
   res, err := http.Post(
      "https://apis.justwatch.com/graphql", "application/json",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   var v struct {
      Data struct {
         URL struct {
            Node struct {
               Offers []Offer
            }
         }
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
      return nil, err
   }
   return v.Data.URL.Node.Offers, nil
}
