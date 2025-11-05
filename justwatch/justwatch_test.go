package justwatch

import (
   "fmt"
   "testing"
)

func Test(t *testing.T) {
   localesVar, err := FetchLocales("en-US")
   if err != nil {
      t.Fatal(err)
   }
   for _, l := range localesVar {
      fmt.Printf("{%q, %q, %q},\n", l.FullLocale, l.Country, l.CountryName)
   }
   localeVar, ok := localesVar.Locale(&HrefLangTag{Locale: "en_US"})
   if !ok {
      t.Fatal("Locales.Locale")
   }
   fmt.Println(localeVar)
}
