package justwatch

import (
   "bytes"
   "cmp"
   "encoding/json"
   "errors"
   "html"
   "net/http"
   "slices"
   "strconv"
   "strings"
)

func (gs OfferGroups) String() string {
   var b []byte
   slices.SortFunc(gs, func(c, d *OfferGroup) int {
      if v := len(d.Country) - len(c.Country); v != 0 {
         return v
      }
      return cmp.Compare(c.Url, d.Url)
   })
   for i, g := range gs {
      if i >= 1 {
         b = append(b, "\n\n"...)
      }
      b = append(b, "url = "...)
      b = append(b, html.UnescapeString(g.Url)...)
      b = append(b, "\nmonetization = "...)
      b = append(b, g.Monetization...)
      if v := g.Count; v >= 1 {
         b = append(b, "\ncount = "...)
         b = strconv.AppendInt(b, v, 10)
      }
      slices.Sort(g.Country)
      for _, country := range g.Country {
         b = append(b, "\ncountry = "...)
         b = append(b, country...)
      }
   }
   return string(b)
}

func (gs *OfferGroups) Add(s *LocaleState, n OfferNode) {
   i := slices.IndexFunc(*gs, func(g *OfferGroup) bool {
      return g.Url == string(n.StandardWebUrl)
   })
   if i >= 0 {
      g := (*gs)[i]
      if !slices.Contains(g.Country, s.CountryName) {
         g.Country = append(g.Country, s.CountryName)
      }
   } else {
      var g OfferGroup
      g.Count = n.ElementCount
      g.Country = []string{s.CountryName}
      g.Monetization = n.MonetizationType
      g.Url = string(n.StandardWebUrl)
      *gs = append(*gs, &g)
   }
}

type OfferGroup struct {
   Count int64
   Country []string
   Monetization string
   Url string
}

type OfferGroups []*OfferGroup

// `presentationType` data seems to be incorrect in some cases. For example,
// JustWatch reports this as SD: fetchtv.com.au/movie/details/19285
// when the site itself reports as HD
type OfferNode struct {
   ElementCount int64
   MonetizationType string
   StandardWebUrl WebUrl
}

func (t LangTag) Offers(state *LocaleState) ([]OfferNode, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Variables struct {
            Country string `json:"country"`
            FullPath string `json:"fullPath"`
         }
         Query string
      }
      s.Query = graphql_compact(title_details)
      s.Variables.FullPath = t.Href
      s.Variables.Country = state.Country
      return json.Marshal(s)
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
   var data struct {
      Data struct {
         Url struct {
            Node struct {
               Offers []OfferNode
            }
         }
      }
   }
   if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
      return nil, err
   }
   return data.Data.Url.Node.Offers, nil
}

const title_details = `
query(
   $fullPath: String!
   $country: Country!
   $platform: Platform! = WEB
) {
   url(fullPath: $fullPath) {
      node {
         ... on MovieOrShowOrSeason {
            offers(country: $country, platform: $platform) {
               elementCount
               monetizationType
               standardWebURL
            }
         }
      }
   }
}
`

var contains []string

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

func Url(o OfferNode) bool {
   for _, value := range contains {
      if strings.Contains(string(o.StandardWebUrl), value) {
         return true
      }
   }
   return false
}

type WebUrl string

func (w *WebUrl) UnmarshalText(text []byte) error {
   text = bytes.TrimSuffix(text, []byte{'\n'})
   *w = WebUrl(text)
   return nil
}

