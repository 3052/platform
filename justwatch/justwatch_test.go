package justwatch

import (
   "fmt"
   "testing"
)

func TestLocale(t *testing.T) {
   var locales LocaleStates
   err := locales.Make("en-US")
   if err != nil {
      t.Fatal(err)
   }
   for _, v := range locales {
      fmt.Printf("{%q, %q, %q},\n", v.FullLocale, v.Country, v.CountryName)
   }
   tag := LangTag{Locale: "en_US"}
   fmt.Println(locales.Locale(tag))
}
