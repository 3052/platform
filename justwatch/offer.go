package justwatch

import (
   "bytes"
   "cmp"
   "encoding/json"
   "errors"
   "html"
   "net/http"
   "slices"
   "strings"
)

var contains []string

func (t LangTag) Offers(s *LocaleState) ([]OfferNode, error) {
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
   resp, err := http.Post(
      "https://apis.justwatch.com/graphql", "application/json",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   var v struct {
      Data struct {
         URL struct {
            Node struct {
               Offers []OfferNode
            }
         }
      }
   }
   if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
      return nil, err
   }
   return v.Data.URL.Node.Offers, nil
}

type OfferGroups []*OfferGroup

const title_details = `
query GetUrlTitleDetails(
   $fullPath: String!
   $country: Country!
   $platform: Platform! = WEB
) {
   url(fullPath: $fullPath) {
      node {
         ... on MovieOrShowOrSeason {
            offers(country: $country, platform: $platform) {
               monetizationType
               standardWebURL
            }
         }
      }
   }
}
`

type OfferGroup struct {
   URL string
   Monetization string
   Country []string
}

func Monetization(o OfferNode) bool {
   switch o.MonetizationType {
   case "BUY":
      return true
   case "CINEMA":
      return true
   case "RENT":
      return true
   }
   return false
}

func URL(o OfferNode) bool {
   for _, value := range contains {
      if strings.Contains(string(o.StandardWebUrl), value) {
         return true
      }
   }
   return false
}

// `presentationType` data seems to be incorrect in some cases. For example,
// JustWatch reports this as SD: fetchtv.com.au/movie/details/19285
// when the site itself reports as HD
type OfferNode struct {
   MonetizationType string
   StandardWebUrl web_url
}

type web_url string

func (w *web_url) UnmarshalText(text []byte) error {
   text = bytes.TrimSuffix(text, []byte{'\n'})
   *w = web_url(text)
   return nil
}
func (gs *OfferGroups) Add(s *LocaleState, n OfferNode) {
   i := slices.IndexFunc(*gs, func(g *OfferGroup) bool {
      return g.URL == string(n.StandardWebUrl)
   })
   if i >= 0 {
      g := (*gs)[i]
      if !slices.Contains(g.Country, s.CountryName) {
         g.Country = append(g.Country, s.CountryName)
      }
   } else {
      var g OfferGroup
      g.URL = string(n.StandardWebUrl)
      g.Monetization = n.MonetizationType
      g.Country = []string{s.CountryName}
      *gs = append(*gs, &g)
   }
}

func (gs OfferGroups) String() string {
   var b strings.Builder
   slices.SortFunc(gs, func(a, b *OfferGroup) int {
      if v := len(b.Country) - len(a.Country); v != 0 {
         return v
      }
      return cmp.Compare(a.URL, b.URL)
   })
   for i, g := range gs {
      if i >= 1 {
         b.WriteString("\n\n")
      }
      b.WriteString("url = ")
      b.WriteString(html.UnescapeString(g.URL))
      b.WriteString("\nmonetization = ")
      b.WriteString(g.Monetization)
      slices.Sort(g.Country)
      for _, country := range g.Country {
         b.WriteString("\ncountry = ")
         b.WriteString(country)
      }
   }
   return b.String()
}
