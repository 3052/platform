package justwatch

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "errors"
   "net/http"
)

func (f *FetcherQuery) New(language string) error {
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
   return json.NewDecoder(res.Body).Decode(f)
}

const fetcher_query = `
query BackendConstantsFetcherQuery($language: Language!) {
   locales {
      countryName(language: $language)
      fullLocale
   }
}
`

type FetcherQuery struct {
   Data struct {
      Locales []struct {
         CountryName string
         FullLocale string
      }
   }
}
