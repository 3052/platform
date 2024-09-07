package justwatch

import (
   "fmt"
   "reflect"
   "testing"
)

func TestLocale(t *testing.T) {
   locales, err := NewLocaleStates("en-US")
   if err != nil {
      t.Fatal(err)
   }
   for _, v := range locales {
      fmt.Printf("{%q, %q, %q},\n", v.FullLocale, v.Country, v.CountryName)
   }
   tag := LangTag{Locale: "en_US"}
   fmt.Println(locales.Locale(&tag))
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
