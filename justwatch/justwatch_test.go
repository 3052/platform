package justwatch

import (
   "fmt"
   "reflect"
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

func TestSize(t *testing.T) {
   size := reflect.TypeOf(&struct{}{}).Size()
   for _, test := range size_tests {
      if reflect.TypeOf(test).Size() > size {
         fmt.Printf("*%T\n", test)
      } else {
         fmt.Printf("%T\n", test)
      }
   }
}

var size_tests = []any{
   Address{},
   ContentUrls{},
   LangTag{},
   LocaleState{},
   LocaleStates{},
   OfferGroup{},
   OfferGroups{},
   OfferNode{},
   WebUrl{},
}
