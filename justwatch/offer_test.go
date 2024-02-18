package justwatch

import (
   "fmt"
   "testing"
   "time"
)

// justwatch.com/us/movie/mulholland-drive
const movie = "/us/movie/mulholland-drive"

func TestContent(t *testing.T) {
   var content ContentUrls
   err := content.New(movie)
   if err != nil {
     t.Fatal(err)
   }
   for i, tag := range content.Href_Lang_Tags {
      if i >= 9 {
         break
      }
      locale, ok := EnglishLocales.Locale(tag)
      if !ok {
         t.Fatal(tag)
      }
      offers, err := tag.Offers(locale)
      if err != nil {
         t.Fatal(err)
      }
      for i, offer := range offers {
         if i >= 1 {
            fmt.Println()
         }
         if offer.Stream() {
            fmt.Println(locale.String(offer))
         }
      }
      time.Sleep(99 * time.Millisecond)
   }
}
