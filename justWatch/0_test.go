package justWatch

import (
   "fmt"
   "testing"
)

func Test(t *testing.T) {
   localesVar, err := FetchLocales("en-US")
   if err != nil {
      t.Fatal(err)
   }
   for _, localeVar := range localesVar {
      fmt.Printf("%#v,\n", localeVar)
   }
   localeVar, ok := localesVar.Locale(&HrefLangTag{Locale: "en_US"})
   if !ok {
      t.Fatal("Locales.Locale")
   }
   fmt.Println(localeVar)
}
