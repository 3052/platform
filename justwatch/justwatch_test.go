package justwatch

import (
   "fmt"
   "testing"
)

func TestLocale(t *testing.T) {
   var locales1 Locales
   err := locales1.New("en-US")
   if err != nil {
      t.Fatal(err)
   }
   for _, v := range locales1 {
      fmt.Printf("{%q, %q, %q},\n", v.FullLocale, v.Country, v.CountryName)
   }
   locale1, ok := locales1.Locale(&LangTag{Locale: "en_US"})
   if !ok {
      t.Fatal("Locales.Locale")
   }
   fmt.Println(locale1)
}
