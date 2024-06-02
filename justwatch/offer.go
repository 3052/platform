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

var contains = []string{
   // 2024-6-2
   "/www.raiplay.it/",
   // 2024-6-1
   "/app.primevideo.com/",
   "/watch.amazon.co.jp/",
   "/watch.amazon.co.uk/",
   "/watch.amazon.com/",
   "/watch.amazon.de/",
   "/www.cineast.fi/",
   "/www.filmfriend.ch/",
   "/www.hulu.jp/",
   // 2024-5-31
   "/hd.kinopoisk.ru/",
   "/hollywoodsuite.ca/",
   "/player.pl/",
   "/skygo.sky.it/",
   "/tv.allente.dk/",
   "/tvplus.com.tr/",
   "/viaplay.dk/",
   "/viaplay.fi/",
   "/viaplay.no/",
   "/viaplay.se/",
   "/web.magentatv.de/",
   "/www.cineast.no/",
   "/www.filmbox.com/",
   "/www.filmotv.fr/",
   "/www.hoopladigital.com/",
   "/www.nowtv.com/",
   "/www.sky.com/",
   "/www.tbs.com/",
   "/www.tntdrama.com/",
   "/www.trutv.com/",
   // 2024-5-29
   "/hbogo.co.th/",
   "/hbogoasia.id/",
   "/hbogoasia.my/",
   "/hbogoasia.ph/",
   "/hbogoasia.sg/",
   "/hbogoasia.tw/",
   "/more.tv/",
   // 2024-5-26
   "/tv.kinopop.com/",
   "/www.canalplus.com/ch/",
   "/www.france.tv/",
   "/www.mycanal.fr/",
   "/www.starz.com/",
   // 2024-5-23
   "/disneyplus.bn5x.net/",
   // 2024-5-16
   "/binge.com.au/",
   "/filmoteket.no/",
   "/fjernleje.filmstriben.dk/",
   "/globoplay.globo.com/",
   "/maxstream.tv/",
   "/okko.tv/",
   "/osnplus.com/",
   "/play.movistar.com.ec/",
   "/premier.one/",
   "/tv.kpn.com/",
   "/video.unext.jp/",
   "/watcha.com/",
   "/wl.movistarplus.es/",
   "/www.canalplus.at/",
   "/www.catchplay.com/",
   "/www.cineasterna.com/",
   "/www.crave.ca/",
   "/www.jiocinema.com/",
   "/www.movistarplay.cl/",
   "/www.movistarplay.co/",
   "/www.neontv.co.nz/",
   "/www.nowtv.it/",
   "/www.play.movistar.com.ar/",
   "/www.starplus.com/",
   "/www.strim.no/",
   "/www.sunrisetv.ch/",
   // 2024-5-15
   "/www.foxtel.com.au/",
   "/www.joyn.de/",
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

