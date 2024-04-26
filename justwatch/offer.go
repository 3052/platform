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
   // 2024-4-25
   "/binge.com.au/",
   "/disneyplus.bn5x.net/",
   "/filmoteket.no/",
   "/fjernleje.filmstriben.dk/",
   "/globoplay.globo.com/",
   "/hbogo.co.th/",
   "/hbogoasia.id/",
   "/hbogoasia.my/",
   "/hbogoasia.ph/",
   "/hbogoasia.sg/",
   "/hbogoasia.tw/",
   "/hd.kinopoisk.ru/",
   "/hollywoodsuite.ca/",
   "/mediasetinfinity.mediaset.it/",
   "/more.tv/",
   "/www.comhemplay.se/",
   "/oiplay.tv/",
   "/okko.tv/",
   "/osnplus.com/",
   "/play.betvgo.be/",
   "/play.movistar.com.ec/",
   "/player.pl/",
   "/premier.one/",
   "/starzplay.com/ar/",
   "/stream.sooner.be/",
   "/tv.allente.dk/",
   "/tv.kpn.com/",
   "/tvplus.com.tr/",
   "/viaplay.dk/",
   "/viaplay.fi/",
   "/viaplay.nl/",
   "/viaplay.no/",
   "/viaplay.se/",
   "/video.unext.jp/",
   "/watcha.com/",
   "/web.magentatv.de/",
   "/wl.movistarplus.es/",
   "/www.blutv.com/",
   "/www.canalplus.at/",
   "/www.canalplus.com/ch/",
   "/www.catchplay.com/id/",
   "/www.catchplay.com/tw/",
   "/www.crave.ca/",
   "/www.filmbox.com/pl/",
   "/www.filmin.es/",
   "/www.foxtel.com.au/",
   "/www.hoopladigital.com/",
   "/www.hotstar.com/id/",
   "/www.hotstar.com/my/",
   "/www.hotstar.com/th/",
   "/www.jiocinema.com/",
   "/www.joyn.de/",
   "/www.movistarplay.cl/",
   "/www.movistarplay.co/",
   "/www.mycanal.fr/",
   "/www.neontv.co.nz/",
   "/www.nowonline.com.br/",
   "/www.nowtv.com/",
   "/www.nowtv.it/",
   "/www.play.movistar.com.ar/",
   "/www.rakuten.tv/dk/",
   "/www.rakuten.tv/fi/",
   "/www.rakuten.tv/fr/",
   "/www.rakuten.tv/no/",
   "/www.rakuten.tv/se/",
   "/www.showmax.com/",
   "/www.sky.com/watch/sky-go/",
   "/www.sonyliv.com/",
   "/www.starplus.com/",
   "/www.strim.no/",
   "/www.sunrisetv.ch/",
   "/www.timvision.it/",
   "/www.tivify.tv/",
   "/www.tod.tv/",
   "/www.wavve.com/",
   "/www.xn--lep-tma39c.tv/",
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

// `presentationType` data seems to be incorrect in some cases. For example,
// JustWatch reports this as SD: fetchtv.com.au/movie/details/19285
// when the site itself reports as HD
type OfferNode struct {
   MonetizationType string
   StandardWebUrl string
}

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
      if strings.Contains(o.StandardWebUrl, value) {
         return true
      }
   }
   return false
}
