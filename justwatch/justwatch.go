package justwatch

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "errors"
   "html"
   "net/http"
   "slices"
   "strconv"
   "strings"
)

func (o Offer) Monetization() bool {
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

type Content struct {
   HrefLangTags []LangTag `json:"href_lang_tags"`
}

type Address struct {
   s string
}

type WebUrl struct {
   S string
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

func (t *LangTag) Offers(locale0 *Locale) ([]Offer, error) {
   var value struct {
      Query string `json:"query"`
      Variables struct {
         Country string `json:"country"`
         FullPath string `json:"fullPath"`
      } `json:"variables"`
   }
   value.Query = graphql_compact(title_details)
   value.Variables.Country = locale0.Country
   value.Variables.FullPath = t.Href
   data, err := json.Marshal(value)
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
      var data strings.Builder
      resp.Write(&data)
      return nil, errors.New(data.String())
   }
   var value1 struct {
      Data struct {
         Url struct {
            Node struct {
               Offers []Offer
            }
         }
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value1)
   if err != nil {
      return nil, err
   }
   return value1.Data.Url.Node.Offers, nil
}

func (s *Locales) New(language string) error {
   var value struct {
      Query     string `json:"query"`
      Variables struct {
         Language string `json:"language"`
      } `json:"variables"`
   }
   value.Query = graphql_compact(fetcher_query)
   value.Variables.Language = language
   data, err := json.Marshal(value)
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://apis.justwatch.com/graphql", bytes.NewReader(data),
   )
   if err != nil {
      return err
   }
   req.Header.Set("content-type", "application/json")
   req.Header.Set(
      "device-id", base64.RawStdEncoding.EncodeToString(make([]byte, 16)),
   )
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var data bytes.Buffer
      resp.Write(&data)
      return errors.New(data.String())
   }
   var value1 struct {
      Data struct {
         Locales Locales
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value1)
   if err != nil {
      return err
   }
   *s = value1.Data.Locales
   return nil
}

var English = Locales{
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

// this is better than strings.Replace and strings.ReplaceAll
func graphql_compact(data string) string {
   return strings.Join(strings.Fields(data), " ")
}

func (a *Address) Set(data string) error {
   a.s = strings.TrimPrefix(data, "https://")
   a.s = strings.TrimPrefix(a.s, "www.")
   a.s = strings.TrimPrefix(a.s, "justwatch.com")
   return nil
}

func (a Address) String() string {
   return a.s
}

func (w *WebUrl) UnmarshalText(data []byte) error {
   w.S = strings.TrimSuffix(string(data), "\n")
   return nil
}

func (a Address) Content() (*Content, error) {
   resp, err := http.Get("https://apis.justwatch.com/content/urls?path=" + a.s)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   content0 := &Content{}
   err = json.NewDecoder(resp.Body).Decode(content0)
   if err != nil {
      return nil, err
   }
   return content0, nil
}

type Locales []Locale

// keep order
type Locale struct {
   FullLocale  string
   Country     string
   CountryName string
}

type LangTag struct {
   Locale string // es_AR
   Href   string // /ar/pelicula/mulholland-drive
}

// `presentationType` data seems to be incorrect in some cases. For example,
// JustWatch reports this as SD: fetchtv.com.au/movie/details/19285
// when the site itself reports as HD
type Offer struct {
   ElementCount     int64
   MonetizationType string
   StandardWebUrl   WebUrl
}

func (s Locales) Locale(tag *LangTag) (*Locale, bool) {
   for _, locale0 := range s {
      if locale0.FullLocale == tag.Locale {
         return &locale0, true
      }
   }
   return nil, false
}

type OfferRow struct {
   Count        int64
   Country      []string
   Monetization string
   Url          string
}

type OfferRows []*OfferRow

func (o *OfferRows) Add(locale0 *Locale, offer0 *Offer) {
   i := slices.IndexFunc(*o, func(row *OfferRow) bool {
      return row.Url == offer0.StandardWebUrl.S
   })
   if i >= 0 {
      row := (*o)[i]
      if !slices.Contains(row.Country, locale0.CountryName) {
         row.Country = append(row.Country, locale0.CountryName)
      }
   } else {
      var row OfferRow
      row.Count = offer0.ElementCount
      row.Country = []string{locale0.CountryName}
      row.Monetization = offer0.MonetizationType
      row.Url = offer0.StandardWebUrl.S
      *o = append(*o, &row)
   }
}

func (o OfferRows) String() string {
   var data []byte
   slices.SortFunc(o, func(a, b *OfferRow) int {
      return len(a.Url) - len(b.Url)
   })
   for i, row := range o {
      if i >= 1 {
         data = append(data, "\n\n"...)
      }
      data = append(data, "url = "...)
      data = append(data, html.UnescapeString(row.Url)...)
      data = append(data, "\nmonetization = "...)
      data = append(data, row.Monetization...)
      if row.Count >= 1 {
         data = append(data, "\ncount = "...)
         data = strconv.AppendInt(data, row.Count, 10)
      }
      slices.Sort(row.Country)
      for _, country := range row.Country {
         data = append(data, "\ncountry = "...)
         data = append(data, country...)
      }
   }
   return string(data)
}
