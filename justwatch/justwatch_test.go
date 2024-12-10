package justwatch

import (
   "fmt"
   "testing"
)

func TestLocale(t *testing.T) {
   var states LocaleStates
   err := states.New("en-US")
   if err != nil {
      t.Fatal(err)
   }
   for _, state := range states {
      fmt.Printf(
         "{%q, %q, %q},\n", state.FullLocale, state.Country, state.CountryName,
      )
   }
   locale, ok := states.Locale(&LangTag{Locale: "en_US"})
   if !ok {
      t.Fatal("LocaleStates.Locale")
   }
   fmt.Println(locale)
}
