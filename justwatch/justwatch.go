package justwatch

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "errors"
   "log/slog"
   "net/http"
   "strings"
)

func (s *LocaleStates) Make(language string) error {
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
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://apis.justwatch.com/graphql", bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   device := base64.RawStdEncoding.EncodeToString(make([]byte, 16))
   req.Header = http.Header{
      "Content-Type": {"application/json"},
      "Device-Id": {device},
   }
   slog.Info(req.Method, "URL", req.URL)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b bytes.Buffer
      res.Write(&b)
      return errors.New(b.String())
   }
   var query struct {
      Data struct {
         Locales LocaleStates
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&query); err != nil {
      return err
   }
   *s = query.Data.Locales
   return nil
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

type LocaleState struct {
   FullLocale string
   Country string
   CountryName string
}

type LocaleStates []LocaleState

func (s LocaleStates) Locale(t LangTag) (*LocaleState, bool) {
   for _, locale := range s {
      if locale.FullLocale == t.Locale {
         return &locale, true
      }
   }
   return nil, false
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

type LangTag struct {
   Href string // /ar/pelicula/mulholland-drive
   Locale string // es_AR
}

// this is better than strings.Replace and strings.ReplaceAll
func graphql_compact(s string) string {
   field := strings.Fields(s)
   return strings.Join(field, " ")
}

type ContentUrls struct {
   Href_Lang_Tags []LangTag
}

func (c *ContentUrls) New(path string) error {
   path = "https://apis.justwatch.com/content/urls?path=" + path
   slog.Info("GET", "URL", path)
   res, err := http.Get(path)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return json.NewDecoder(res.Body).Decode(c)
}
