package justwatch

import (
   "fmt"
   "testing"
)

func TestLocale(t *testing.T) {
   var localesVar Locales
   err := localesVar.New("en-US")
   if err != nil {
      t.Fatal(err)
   }
   for _, v := range localesVar {
      fmt.Printf("{%q, %q, %q},\n", v.FullLocale, v.Country, v.CountryName)
   }
   localeVar, ok := localesVar.Locale(&LangTag{Locale: "en_US"})
   if !ok {
      t.Fatal("Locales.Locale")
   }
   fmt.Println(localeVar)
}
