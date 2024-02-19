package justwatch

import (
   "fmt"
   "slices"
   "testing"
   "time"
)

// justwatch.com/us/movie/mulholland-drive
const movie = "/us/movie/mulholland-drive"

func TestOffer(t *testing.T) {
   var content ContentUrls
   err := content.New(movie)
   if err != nil {
     t.Fatal(err)
   }
   var groups OfferGroups
   for _, tag := range content.Href_Lang_Tags {
      fmt.Printf("%+v\n", tag)
      locale, ok := EnglishLocales.Locale(tag)
      if !ok {
         t.Fatal(tag)
      }
      offers, err := tag.Offers(locale)
      if err != nil {
         t.Fatal(err)
      }
      for _, offer := range slices.DeleteFunc(offers, Delete) {
         groups.Add(locale, offer)
      }
      time.Sleep(99 * time.Millisecond)
   }
   fmt.Println(groups)
}
