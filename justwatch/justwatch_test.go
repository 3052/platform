package justwatch

import (
   "cmp"
   "encoding/json"
   "fmt"
   "log"
   "maps"
   "slices"
   "testing"
)

func Test(t *testing.T) {
   var localesVar Locales
   err := localesVar.New("en-US")
   if err != nil {
      t.Fatal(err)
   }
   for _, l := range localesVar {
      fmt.Printf("{%q, %q, %q},\n", l.FullLocale, l.Country, l.CountryName)
   }
   localeVar, ok := localesVar.Locale(&LangTag{Locale: "en_US"})
   if !ok {
      t.Fatal("Locales.Locale")
   }
   fmt.Println(localeVar)
}
