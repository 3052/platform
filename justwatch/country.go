package justwatch

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "errors"
   "net/http"
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
      countryName(language: $language)
      fullLocale
   }
}
`

type LocaleStates []struct {
   FullLocale string
   CountryName string
}

func (s LocaleStates) Country(locale string) (string, bool) {
   for _, each := range s {
      if each.FullLocale == locale {
         return each.CountryName, true
      }
   }
   return "", false
}

var EnglishLocales = LocaleStates{
   {"ar_AE", "United Arab Emirates"},
   {"ar_BH", "Bahrain"},
   {"ar_DZ", "Algeria"},
   {"ar_EG", "Egypt"},
   {"ar_IQ", "Iraq"},
   {"ar_JO", "Jordan"},
   {"ar_KW", "Kuwait"},
   {"ar_LB", "Lebanon"},
   {"ar_LY", "Libya"},
   {"ar_MA", "Morocco"},
   {"ar_OM", "Oman"},
   {"ar_PS", "Palestine"},
   {"ar_QA", "Qatar"},
   {"ar_SA", "Saudi Arabia"},
   {"ar_TD", "Chad"},
   {"ar_TN", "Tunisia"},
   {"ar_YE", "Yemen"},
   {"az_AZ", "Azerbaijan"},
   {"be_BY", "Belarus"},
   {"bg_BG", "Bulgaria"},
   {"bs_BA", "Bosnia and Herzegovina"},
   {"ca_AD", "Andorra"},
   {"cs_CZ", "Czech Republic"},
   {"de_AT", "Austria"},
   {"de_CH", "Switzerland"},
   {"de_DE", "Germany"},
   {"de_LI", "Liechtenstein"},
   {"el_CY", "Cyprus"},
   {"el_GR", "Greece"},
   {"en_AG", "Antigua and Barbuda"},
   {"en_AU", "Australia"},
   {"en_BB", "Barbados"},
   {"en_BM", "Bermuda"},
   {"en_BS", "Bahamas"},
   {"en_BZ", "Belize"},
   {"en_CA", "Canada"},
   {"en_CM", "Cameroon"},
   {"en_DK", "Denmark"},
   {"en_FJ", "Fiji"},
   {"en_GB", "United Kingdom"},
   {"en_GG", "Guernsey"},
   {"en_GH", "Ghana"},
   {"en_GI", "Gibraltar"},
   {"en_GY", "Guyana"},
   {"en_ID", "Indonesia"},
   {"en_IE", "Ireland"},
   {"en_IN", "India"},
   {"en_JM", "Jamaica"},
   {"en_KE", "Kenya"},
   {"en_LC", "Saint Lucia"},
   {"en_MW", "Malawi"},
   {"en_MY", "Malaysia"},
   {"en_NG", "Nigeria"},
   {"en_NL", "Netherlands"},
   {"en_NO", "Norway"},
   {"en_NZ", "New Zealand"},
   {"en_PG", "Papua New Guinea"},
   {"en_PH", "Philippines"},
   {"en_SG", "Singapore"},
   {"en_TC", "Turks and Caicos Islands"},
   {"en_TH", "Thailand"},
   {"en_TT", "Trinidad and Tobago"},
   {"en_UG", "Uganda"},
   {"en_US", "United States"},
   {"en_ZA", "South Africa"},
   {"en_ZM", "Zambia"},
   {"en_ZW", "Zimbabwe"},
   {"es_AR", "Argentina"},
   {"es_BO", "Bolivia"},
   {"es_CL", "Chile"},
   {"es_CO", "Colombia"},
   {"es_CR", "Costa Rica"},
   {"es_CU", "Cuba"},
   {"es_DO", "Dominican Republic"},
   {"es_EC", "Ecuador"},
   {"es_ES", "Spain"},
   {"es_GQ", "Equatorial Guinea"},
   {"es_GT", "Guatemala"},
   {"es_HN", "Honduras"},
   {"es_MX", "Mexico"},
   {"es_NI", "Nicaragua"},
   {"es_PA", "Panama"},
   {"es_PE", "Peru"},
   {"es_PY", "Paraguay"},
   {"es_SV", "El Salvador"},
   {"es_UY", "Uruguay"},
   {"es_VE", "Venezuela"},
   {"et_EE", "Estonia"},
   {"fi_FI", "Finland"},
   {"fr_BE", "Belgium"},
   {"fr_BF", "Burkina Faso"},
   {"fr_CD", "DR Congo"},
   {"fr_CI", "Ivory Coast"},
   {"fr_FR", "France"},
   {"fr_GF", "French Guiana"},
   {"fr_LU", "Luxembourg"},
   {"fr_MC", "Monaco"},
   {"fr_MG", "Madagascar"},
   {"fr_ML", "Mali"},
   {"fr_MU", "Mauritius"},
   {"fr_NE", "Niger"},
   {"fr_PF", "French Polynesia"},
   {"fr_SC", "Seychelles"},
   {"fr_SN", "Senegal"},
   {"he_IL", "Israel"},
   {"hr_HR", "Croatia"},
   {"hu_HU", "Hungary"},
   {"is_IS", "Iceland"},
   {"it_IT", "Italy"},
   {"it_SM", "San Marino"},
   {"it_VA", "Vatican City"},
   {"ja_JP", "Japan"},
   {"ko_KR", "South Korea"},
   {"lt_LT", "Lithuania"},
   {"lv_LV", "Latvia"},
   {"mk_MK", "Macedonia"},
   {"mt_MT", "Malta"},
   {"pl_PL", "Poland"},
   {"pt_AO", "Angola"},
   {"pt_BR", "Brazil"},
   {"pt_CV", "Cape Verde"},
   {"pt_MZ", "Mozambique"},
   {"pt_PT", "Portugal"},
   {"ro_MD", "Moldova"},
   {"ro_RO", "Romania"},
   {"ru_RU", "Russia"},
   {"sk_SK", "Slovakia"},
   {"sl_SI", "Slovenia"},
   {"sq_AL", "Albania"},
   {"sq_XK", "Kosovo"},
   {"sr_ME", "Montenegro"},
   {"sr_RS", "Serbia"},
   {"sv_SE", "Sweden"},
   {"sw_TZ", "Tanzania"},
   {"tr_TR", "Turkey"},
   {"uk_UA", "Ukraine"},
   {"ur_PK", "Pakistan"},
   {"zh_HK", "Hong Kong"},
   {"zh_TW", "Taiwan"},
}
