package justwatch

import (
   "bytes"
   "cmp"
   "encoding/base64"
   "encoding/json"
   "errors"
   "html"
   "net/http"
   "slices"
   "strconv"
   "strings"
)

var contains = []string{
   // 2024-7-20
   "/www.stan.com.au/",
}

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
   var body struct {
      Variables struct {
         Country string `json:"country"`
         FullPath string `json:"fullPath"`
      }
      Query string
   }
   body.Query = graphql_compact(title_details)
   body.Variables.FullPath = t.Href
   body.Variables.Country = state.Country
   data, err := json.Marshal(body)
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
   var resp_body struct {
      Data struct {
         Url struct {
            Node struct {
               Offers []OfferNode
            }
         }
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&resp_body)
   if err != nil {
      return nil, err
   }
   return resp_body.Data.Url.Node.Offers, nil
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

func NewLocaleStates(language string) (LocaleStates, error) {
   body, err := func() ([]byte, error) {
      var s struct {
         Query string
         Variables struct {
            Language string `json:"language"`
         }
      }
      s.Query = graphql_compact(fetcher_query)
      s.Variables.Language = language
      return json.Marshal(s)
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://apis.justwatch.com/graphql", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   device := base64.RawStdEncoding.EncodeToString(make([]byte, 16))
   req.Header = http.Header{
      "content-type": {"application/json"},
      "device-id": {device},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b bytes.Buffer
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   var data struct {
      Data struct {
         Locales LocaleStates
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&data)
   if err != nil {
      return nil, err
   }
   return data.Data.Locales, nil
}

func (LocaleState) Error() string {
   return "LocaleState"
}

// keep order
type LocaleState struct {
   FullLocale string
   Country string
   CountryName string
}

var EnglishLocales = LocaleStates{
   {"ar_AE", "AE", "United Arab Emirates"},
   {"ar_BH", "BH", "Bahrain"},
   {"ar_DZ", "DZ", "Algeria"},
   {"ar_EG", "EG", "Egypt"},
   {"ar_IQ", "IQ", "Iraq"},
   {"ar_JO", "JO", "Jordan"},
   {"ar_KW", "KW", "Kuwait"},
   {"ar_LB", "LB", "Lebanon"},
   {"ar_LY", "LY", "Libya"},
   {"ar_MA", "MA", "Morocco"},
   {"ar_OM", "OM", "Oman"},
   {"ar_PS", "PS", "Palestine"},
   {"ar_QA", "QA", "Qatar"},
   {"ar_SA", "SA", "Saudi Arabia"},
   {"ar_TD", "TD", "Chad"},
   {"ar_TN", "TN", "Tunisia"},
   {"ar_YE", "YE", "Yemen"},
   {"az_AZ", "AZ", "Azerbaijan"},
   {"be_BY", "BY", "Belarus"},
   {"bg_BG", "BG", "Bulgaria"},
   {"bs_BA", "BA", "Bosnia and Herzegovina"},
   {"ca_AD", "AD", "Andorra"},
   {"cs_CZ", "CZ", "Czech Republic"},
   {"de_AT", "AT", "Austria"},
   {"de_CH", "CH", "Switzerland"},
   {"de_DE", "DE", "Germany"},
   {"de_LI", "LI", "Liechtenstein"},
   {"el_CY", "CY", "Cyprus"},
   {"el_GR", "GR", "Greece"},
   {"en_AG", "AG", "Antigua and Barbuda"},
   {"en_AU", "AU", "Australia"},
   {"en_BB", "BB", "Barbados"},
   {"en_BM", "BM", "Bermuda"},
   {"en_BS", "BS", "Bahamas"},
   {"en_BZ", "BZ", "Belize"},
   {"en_CA", "CA", "Canada"},
   {"en_CM", "CM", "Cameroon"},
   {"en_DK", "DK", "Denmark"},
   {"en_FJ", "FJ", "Fiji"},
   {"en_GB", "GB", "United Kingdom"},
   {"en_GG", "GG", "Guernsey"},
   {"en_GH", "GH", "Ghana"},
   {"en_GI", "GI", "Gibraltar"},
   {"en_GY", "GY", "Guyana"},
   {"en_ID", "ID", "Indonesia"},
   {"en_IE", "IE", "Ireland"},
   {"en_IN", "IN", "India"},
   {"en_JM", "JM", "Jamaica"},
   {"en_KE", "KE", "Kenya"},
   {"en_LC", "LC", "Saint Lucia"},
   {"en_MW", "MW", "Malawi"},
   {"en_MY", "MY", "Malaysia"},
   {"en_NG", "NG", "Nigeria"},
   {"en_NL", "NL", "Netherlands"},
   {"en_NO", "NO", "Norway"},
   {"en_NZ", "NZ", "New Zealand"},
   {"en_PG", "PG", "Papua New Guinea"},
   {"en_PH", "PH", "Philippines"},
   {"en_SG", "SG", "Singapore"},
   {"en_TC", "TC", "Turks and Caicos Islands"},
   {"en_TH", "TH", "Thailand"},
   {"en_TT", "TT", "Trinidad and Tobago"},
   {"en_UG", "UG", "Uganda"},
   {"en_US", "US", "United States"},
   {"en_ZA", "ZA", "South Africa"},
   {"en_ZM", "ZM", "Zambia"},
   {"en_ZW", "ZW", "Zimbabwe"},
   {"es_AR", "AR", "Argentina"},
   {"es_BO", "BO", "Bolivia"},
   {"es_CL", "CL", "Chile"},
   {"es_CO", "CO", "Colombia"},
   {"es_CR", "CR", "Costa Rica"},
   {"es_CU", "CU", "Cuba"},
   {"es_DO", "DO", "Dominican Republic"},
   {"es_EC", "EC", "Ecuador"},
   {"es_ES", "ES", "Spain"},
   {"es_GQ", "GQ", "Equatorial Guinea"},
   {"es_GT", "GT", "Guatemala"},
   {"es_HN", "HN", "Honduras"},
   {"es_MX", "MX", "Mexico"},
   {"es_NI", "NI", "Nicaragua"},
   {"es_PA", "PA", "Panama"},
   {"es_PE", "PE", "Peru"},
   {"es_PY", "PY", "Paraguay"},
   {"es_SV", "SV", "El Salvador"},
   {"es_UY", "UY", "Uruguay"},
   {"es_VE", "VE", "Venezuela"},
   {"et_EE", "EE", "Estonia"},
   {"fi_FI", "FI", "Finland"},
   {"fr_BE", "BE", "Belgium"},
   {"fr_BF", "BF", "Burkina Faso"},
   {"fr_CD", "CD", "DR Congo"},
   {"fr_CI", "CI", "Ivory Coast"},
   {"fr_FR", "FR", "France"},
   {"fr_GF", "GF", "French Guiana"},
   {"fr_LU", "LU", "Luxembourg"},
   {"fr_MC", "MC", "Monaco"},
   {"fr_MG", "MG", "Madagascar"},
   {"fr_ML", "ML", "Mali"},
   {"fr_MU", "MU", "Mauritius"},
   {"fr_NE", "NE", "Niger"},
   {"fr_PF", "PF", "French Polynesia"},
   {"fr_SC", "SC", "Seychelles"},
   {"fr_SN", "SN", "Senegal"},
   {"he_IL", "IL", "Israel"},
   {"hr_HR", "HR", "Croatia"},
   {"hu_HU", "HU", "Hungary"},
   {"is_IS", "IS", "Iceland"},
   {"it_IT", "IT", "Italy"},
   {"it_SM", "SM", "San Marino"},
   {"it_VA", "VA", "Vatican City"},
   {"ja_JP", "JP", "Japan"},
   {"ko_KR", "KR", "South Korea"},
   {"lt_LT", "LT", "Lithuania"},
   {"lv_LV", "LV", "Latvia"},
   {"mk_MK", "MK", "Macedonia"},
   {"mt_MT", "MT", "Malta"},
   {"pl_PL", "PL", "Poland"},
   {"pt_AO", "AO", "Angola"},
   {"pt_BR", "BR", "Brazil"},
   {"pt_CV", "CV", "Cape Verde"},
   {"pt_MZ", "MZ", "Mozambique"},
   {"pt_PT", "PT", "Portugal"},
   {"ro_MD", "MD", "Moldova"},
   {"ro_RO", "RO", "Romania"},
   {"ru_RU", "RU", "Russia"},
   {"sk_SK", "SK", "Slovakia"},
   {"sl_SI", "SI", "Slovenia"},
   {"sq_AL", "AL", "Albania"},
   {"sq_XK", "XK", "Kosovo"},
   {"sr_ME", "ME", "Montenegro"},
   {"sr_RS", "RS", "Serbia"},
   {"sv_SE", "SE", "Sweden"},
   {"sw_TZ", "TZ", "Tanzania"},
   {"tr_TR", "TR", "Turkey"},
   {"uk_UA", "UA", "Ukraine"},
   {"ur_PK", "PK", "Pakistan"},
   {"zh_HK", "HK", "Hong Kong"},
   {"zh_TW", "TW", "Taiwan"},
}

const fetcher_query = `
query BackendConstantsFetcherQuery($language: Language!) {
   locales {
      country
      countryName(language: $language)
      fullLocale
   }
}
`

// this is better than strings.Replace and strings.ReplaceAll
func graphql_compact(s string) string {
   field := strings.Fields(s)
   return strings.Join(field, " ")
}

type ContentUrls struct {
   HrefLangTags []LangTag `json:"href_lang_tags"`
}

type LangTag struct {
   Locale string // es_AR
   Href string // /ar/pelicula/mulholland-drive
}

func (a Address) String() string {
   return a.Path
}

type Address struct {
   Path string
}

func (a *Address) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "www.")
   a.Path = strings.TrimPrefix(s, "justwatch.com")
   return nil
}

func (a Address) Content() (*ContentUrls, error) {
   resp, err := http.Get(
      "https://apis.justwatch.com/content/urls?path=" + a.Path,
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   content := &ContentUrls{}
   err = json.NewDecoder(resp.Body).Decode(content)
   if err != nil {
      return nil, err
   }
   return content, nil
}

func (s LocaleStates) Locale(t LangTag) (*LocaleState, bool) {
   for _, locale := range s {
      if locale.FullLocale == t.Locale {
         return &locale, true
      }
   }
   return nil, false
}

type LocaleStates []LocaleState
