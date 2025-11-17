package main

import (
   "41.neocities.org/service/justWatch"
   "errors"
   "flag"
   "fmt"
   "log"
   "net/http"
   "os"
   "path"
   "time"
)

func (f *flag_set) do_address() error {
   url_path, err := justWatch.GetPath(f.address)
   if err != nil {
      return err
   }
   var content justWatch.Content
   err = content.Fetch(url_path)
   if err != nil {
      return err
   }
   var allEnrichedOffers []justWatch.EnrichedOffer
   for _, tag := range content.HrefLangTags {
      locale, ok := justWatch.EnUs.Locale(&tag)
      if !ok {
         return errors.New("Locale")
      }
      log.Print(locale)
      offers, err := tag.Offers(locale)
      if err != nil {
         return err
      }
      for _, offer := range offers {
         allEnrichedOffers = append(allEnrichedOffers,
            justWatch.EnrichedOffer{Offer: offer, Locale: locale},
         )
      }
      time.Sleep(f.sleep)
   }
   enrichedOffers := justWatch.Deduplicate(allEnrichedOffers)
   enrichedOffers = justWatch.FilterOffers(
      enrichedOffers, "BUY", "CINEMA", "RENT",
   )
   sortedUrls, groupedOffers := justWatch.GroupAndSortByURL(enrichedOffers)
   var data []byte
   for i, url := range sortedUrls {
      if i >= 1 {
         data = append(data, '\n')
      }
      data = fmt.Appendln(data, url)
      for i, v := range groupedOffers[url] {
         if i >= 1 {
            data = append(data, '\n')
         }
         data = fmt.Appendln(data, "country =", v.Locale.Country)
         data = fmt.Appendln(data, "name =", v.Locale.CountryName)
         data = fmt.Appendln(data, "monetization =", v.Offer.MonetizationType)
         if v.Offer.ElementCount >= 1 {
            data = fmt.Appendln(data, "count =", v.Offer.ElementCount)
         }
      }
   }
   return write_file(path.Base(url_path), data)
}

func write_file(name string, data []byte) error {
   log.Println("WriteFile", name)
   return os.WriteFile(name, data, os.ModePerm)
}

type flag_set struct {
   address string
   sleep   time.Duration
}

func main() {
   http.DefaultClient.Transport = &justWatch.Transport
   log.SetFlags(log.Ltime)
   var set flag_set
   flag.StringVar(&set.address, "a", "", "address")
   flag.DurationVar(&set.sleep, "s", 99*time.Millisecond, "sleep")
   flag.Parse()
   if set.address != "" {
      err := set.do_address()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
