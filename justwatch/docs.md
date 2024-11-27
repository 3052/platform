# Overview

package `justwatch`

## Index

- [Variables](#variables)
- [Functions](#functions)
  - [func Monetization(node OfferNode) bool](#func-monetization)
  - [func Url(node OfferNode) bool](#func-url)
- [Types](#types)
  - [type Address](#type-address)
    - [func (a \*Address) Content() (\*ContentUrls, error)](#func-address-content)
    - [func (a \*Address) Set(s string) error](#func-address-set)
    - [func (a \*Address) String() string](#func-address-string)
  - [type ContentUrls](#type-contenturls)
  - [type LangTag](#type-langtag)
    - [func (t \*LangTag) Offers(state \*LocaleState) ([]OfferNode, error)](#func-langtag-offers)
  - [type LocaleState](#type-localestate)
  - [type LocaleStates](#type-localestates)
    - [func (s LocaleStates) Locale(tag \*LangTag) (\*LocaleState, bool)](#func-localestates-locale)
    - [func (s \*LocaleStates) New(language string) error](#func-localestates-new)
  - [type OfferGroup](#type-offergroup)
  - [type OfferGroups](#type-offergroups)
    - [func (o \*OfferGroups) Add(node \*OfferNode, state \*LocaleState)](#func-offergroups-add)
    - [func (o OfferGroups) String() string](#func-offergroups-string)
  - [type OfferNode](#type-offernode)
  - [type WebUrl](#type-weburl)
    - [func (w \*WebUrl) UnmarshalText(text []byte) error](#func-weburl-unmarshaltext)
- [Source files](#source-files)

## Variables

```go
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
```

## Functions

### func [Monetization](./justwatch.go#L67)

```go
func Monetization(node OfferNode) bool
```

### func [Url](./justwatch.go#L79)

```go
func Url(node OfferNode) bool
```

## Types

### type [Address](./justwatch.go#L94)

```go
type Address struct {
  Path string
}
```

### func (\*Address) [Content](./justwatch.go#L109)

```go
func (a *Address) Content() (*ContentUrls, error)
```

### func (\*Address) [Set](./justwatch.go#L98)

```go
func (a *Address) Set(s string) error
```

### func (\*Address) [String](./justwatch.go#L105)

```go
func (a *Address) String() string
```

### type [ContentUrls](./justwatch.go#L128)

```go
type ContentUrls struct {
  HrefLangTags []LangTag `json:"href_lang_tags"`
}
```

### type [LangTag](./justwatch.go#L132)

```go
type LangTag struct {
  Locale string // es_AR
  Href   string // /ar/pelicula/mulholland-drive
}
```

### func (\*LangTag) [Offers](./justwatch.go#L137)

```go
func (t *LangTag) Offers(state *LocaleState) ([]OfferNode, error)
```

### type [LocaleState](./justwatch.go#L182)

```go
type LocaleState struct {
  FullLocale  string
  Country     string
  CountryName string
}
```

keep order

### type [LocaleStates](./justwatch.go#L235)

```go
type LocaleStates []LocaleState
```

### func (LocaleStates) [Locale](./justwatch.go#L27)

```go
func (s LocaleStates) Locale(tag *LangTag) (*LocaleState, bool)
```

### func (\*LocaleStates) [New](./justwatch.go#L188)

```go
func (s *LocaleStates) New(language string) error
```

### type [OfferGroup](./justwatch.go#L379)

```go
type OfferGroup struct {
  Count        int64
  Country      []string
  Monetization string
  Url          string
}
```

### type [OfferGroups](./justwatch.go#L434)

```go
type OfferGroups []*OfferGroup
```

### func (\*OfferGroups) [Add](./justwatch.go#L386)

```go
func (o *OfferGroups) Add(node *OfferNode, state *LocaleState)
```

### func (OfferGroups) [String](./justwatch.go#L405)

```go
func (o OfferGroups) String() string
```

### type [OfferNode](./justwatch.go#L439)

```go
type OfferNode struct {
  ElementCount     int64
  MonetizationType string
  StandardWebUrl   WebUrl
}
```

`presentationType` data seems to be incorrect in some cases. For example,
JustWatch reports this as SD: fetchtv.com.au/movie/details/19285
when the site itself reports as HD

### type [WebUrl](./justwatch.go#L445)

```go
type WebUrl struct {
  String string
}
```

### func (\*WebUrl) [UnmarshalText](./justwatch.go#L449)

```go
func (w *WebUrl) UnmarshalText(text []byte) error
```

## Source files

[justwatch.go](./justwatch.go)
