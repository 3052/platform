package justWatch

import (
   "fmt"
   "testing"
)

func Test(t *testing.T) {
   locales_var, err := FetchLocales("en-US")
   if err != nil {
      t.Fatal(err)
   }
   for _, locale_var := range locales_var {
      fmt.Printf("%#v,\n", locale_var)
   }
   locale_var, ok := locales_var.Locale(&HrefLangTag{Locale: "en_US"})
   if !ok {
      t.Fatal("Locales.Locale")
   }
   fmt.Println(locale_var)
}
