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

type WebUrl struct {
   Data string
}

var hosts = []string{
   // 2024-11-2
   "/play.tv2.no/",
   // 2024-11-1
   "/filmoteket.no/",
   "/www.3cat.cat/",
   "/www.cineast.no/",
   // 2024-7-20
   "/www.stan.com.au/",
}

func (s LocaleStates) Locale(tag *LangTag) (*LocaleState, bool) {
   for _, locale := range s {
      if locale.FullLocale == tag.Locale {
         return &locale, true
      }
   }
   return nil, false
}

// NO ANONYMOUS QUERY
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
               elementCount
               monetizationType
               standardWebURL
            }
         }
      }
   }
}
`

const fetcher_query = `
query BackendConstantsFetcherQuery($language: Language!) {
   locales {
      country
      countryName(language: $language)
      fullLocale
   }
}
`

func Monetization(node OfferNode) bool {
   switch node.MonetizationType {
   case "BUY":
      return true
   case "CINEMA":
      return true
   case "RENT":
      return true
   }
   return false
}

func Url(node OfferNode) bool {
   for _, host := range hosts {
      if strings.Contains(node.StandardWebUrl.Data, host) {
         return true
      }
   }
   return false
}

// this is better than strings.Replace and strings.ReplaceAll
func graphql_compact(s string) string {
   field := strings.Fields(s)
   return strings.Join(field, " ")
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

func (a *Address) String() string {
   return a.Path
}

func (a *Address) Content() (*ContentUrls, error) {
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

type ContentUrls struct {
   HrefLangTags []LangTag `json:"href_lang_tags"`
}

type LangTag struct {
   Locale string // es_AR
   Href string // /ar/pelicula/mulholland-drive
}

func (t *LangTag) Offers(state *LocaleState) ([]OfferNode, error) {
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
      bytes.NewReader(data),
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

// keep order
type LocaleState struct {
   FullLocale string
   Country string
   CountryName string
}

func (s *LocaleStates) New(language string) error {
   var body struct {
      Query string
      Variables struct {
         Language string `json:"language"`
      }
   }
   body.Query = graphql_compact(fetcher_query)
   body.Variables.Language = language
   data, err := json.Marshal(body)
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://apis.justwatch.com/graphql", bytes.NewReader(data),
   )
   if err != nil {
      return err
   }
   device := base64.RawStdEncoding.EncodeToString(make([]byte, 16))
   req.Header = http.Header{
      "content-type": {"application/json"},
      "device-id": {device},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b bytes.Buffer
      resp.Write(&b)
      return errors.New(b.String())
   }
   var resp_body struct {
      Data struct {
         Locales LocaleStates
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&resp_body)
   if err != nil {
      return err
   }
   *s = resp_body.Data.Locales
   return nil
}

type LocaleStates []LocaleState

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

type OfferGroup struct {
   Count int64
   Country []string
   Monetization string
   Url string
}

func (o *OfferGroups) Add(node *OfferNode, state *LocaleState) {
   i := slices.IndexFunc(*o, func(group *OfferGroup) bool {
      return group.Url == node.StandardWebUrl.Data
   })
   if i >= 0 {
      group := (*o)[i]
      if !slices.Contains(group.Country, state.CountryName) {
         group.Country = append(group.Country, state.CountryName)
      }
   } else {
      var group OfferGroup
      group.Count = node.ElementCount
      group.Country = []string{state.CountryName}
      group.Monetization = node.MonetizationType
      group.Url = node.StandardWebUrl.Data
      *o = append(*o, &group)
   }
}

func (o OfferGroups) String() string {
   var b []byte
   slices.SortFunc(o, func(c, d *OfferGroup) int {
      if v := len(d.Country) - len(c.Country); v != 0 {
         return v
      }
      return cmp.Compare(c.Url, d.Url)
   })
   for i, group := range o {
      if i >= 1 {
         b = append(b, "\n\n"...)
      }
      b = append(b, "url = "...)
      b = append(b, html.UnescapeString(group.Url)...)
      b = append(b, "\nmonetization = "...)
      b = append(b, group.Monetization...)
      if v := group.Count; v >= 1 {
         b = append(b, "\ncount = "...)
         b = strconv.AppendInt(b, v, 10)
      }
      slices.Sort(group.Country)
      for _, country := range group.Country {
         b = append(b, "\ncountry = "...)
         b = append(b, country...)
      }
   }
   return string(b)
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

func (w *WebUrl) UnmarshalText(text []byte) error {
   w.Data = strings.TrimSuffix(string(text), "\n")
   return nil
}
