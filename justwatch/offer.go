package justwatch

import (
   "bytes"
   "cmp"
   "encoding/json"
   "errors"
   "net/http"
   "slices"
   "strings"
)

func Delete(o OfferNode) bool {
   switch {
   case o.MonetizationType == "BUY":
      return true
   case o.MonetizationType == "RENT":
      return true
   case strings.Contains(o.StandardWebUrl, "/more.tv/"):
      return true
   case strings.Contains(o.StandardWebUrl, "/viddla.fi/"):
      return true
   case strings.Contains(o.StandardWebUrl, "/www.hulu.jp/"):
      return true
   case strings.HasSuffix(o.StandardWebUrl, "/tv.apple.com"):
      return true
   case strings.HasSuffix(o.StandardWebUrl, "/tv.apple.com/de"):
      return true
   }
   return false
}

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
               Offers []OfferNode
            }
         }
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
      return nil, err
   }
   return v.Data.URL.Node.Offers, nil
}

type OfferGroups []*OfferGroup

func (gs *OfferGroups) Add(s *LocaleState, n OfferNode) {
   i := slices.IndexFunc(*gs, func(g *OfferGroup) bool {
      return g.URL == n.StandardWebUrl
   })
   if i >= 0 {
      g := (*gs)[i]
      if !slices.Contains(g.Country, s.CountryName) {
         g.Country = append(g.Country, s.CountryName)
      }
   } else {
      var g OfferGroup
      g.URL = n.StandardWebUrl
      g.Monetization = n.MonetizationType
      g.Country = []string{s.CountryName}
      *gs = append(*gs, &g)
   }
}

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
      b.WriteString(g.URL)
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

// `presentationType` data seems to be incorrect in some cases. For example,
// JustWatch reports this as SD: fetchtv.com.au/movie/details/19285
// when the site itself reports as HD
type OfferNode struct {
   MonetizationType string
   StandardWebUrl string
}

