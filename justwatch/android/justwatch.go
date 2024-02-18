package main

import (
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
   
   "encoding/base64"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "apis.justwatch.com"
   req.URL.Path = "/graphql"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Device-Id"] = []string{
      base64.RawStdEncoding.EncodeToString(make([]byte, 16)),
   }
   body := fmt.Sprintf(`{
    "query": %q,
    "variables": {
     "country": "US",
     "language": "en-US",
     "loggedOut": true,
     "platform": "ANDROID"
    }
   }
   `, fetcher_query)
   req.Body = io.NopCloser(strings.NewReader(body))
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   fmt.Println(string(text))
   if strings.Contains(string(text), `"Argentina"`) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}

const fetcher_query = `
query BackendConstantsFetcherQuery(
  $country: Country!
  $language: Language!
  $loggedOut: Boolean!
  $platform: Platform!
) {
  movieAgeCertifications: ageCertifications(country: $country, objectType: MOVIE) {
    technicalName
  }
  showAgeCertifications: ageCertifications(country: $country, objectType: SHOW) {
    technicalName
  }
  genericTitleLists(type: CUSTOM_LISTS, first: 20) @skip(if: $loggedOut) {
    edges {
      node {
        name
        id
      }
    }
    totalCount
  }
  genres {
    shortName
    technicalName
    id
  }
  languages {
    i18nState
    language
  }
  locales {
    country
    countryName(language: $language)
    currency
    fullLocale
    i18nState
  }
  packages(country: $country, platform: $platform) {
    addons(country: $country, platform: $platform) {
      clearName
      hasSport(country: $country, platform: $platform)
      hasTitles(country: $country, platform: $platform)
      icon
      id
      monetizationTypes
      packageId
      selected
      shortName
      technicalName
    }
    clearName
    hasSport(country: $country, platform: $platform)
    hasTitles(country: $country, platform: $platform)
    icon
    id
    monetizationTypes
    packageId
    selected
    shortName
    technicalName
  }
  stats {
    knowsAboutTitleLists
    maxCustomListsDefault
    maxCustomListsPro
  }
}
`
