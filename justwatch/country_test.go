package justwatch

import (
   "fmt"
   "testing"
)

func TestFetcherQuery(t *testing.T) {
   var query FetcherQuery
   err := query.New("en-US")
   if err != nil {
      t.Fatal(err)
   }
   for _, locale := range query.Data.Locales {
      fmt.Printf("%+v\n", locale)
   }
}
