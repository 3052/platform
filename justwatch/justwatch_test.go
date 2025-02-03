package justwatch

import (
   "fmt"
   "testing"
)

func TestLocale(t *testing.T) {
   var locales0 Locales
   err := locales0.New("en-US")
   if err != nil {
      t.Fatal(err)
   }
   for _, v := range locales0 {
      fmt.Printf("{%q, %q, %q},\n", v.FullLocale, v.Country, v.CountryName)
   }
   locale0, ok := locales0.Locale(&LangTag{Locale: "en_US"})
   if !ok {
      t.Fatal("Locales.Locale")
   }
   fmt.Println(locale0)
}
