package main

import (
   "41.neocities.org/platform/justwatch"
   "errors"
   "flag"
   "fmt"
   "log"
   "net/http"
   "os"
   "path"
   "slices"
   "strings"
   "time"
)

func (f *flag_set) do_address() error {
   url_path, err := justwatch.GetPath(f.address)
   if err != nil {
      return err
   }
   var content justwatch.Content
   err = content.Fetch(url_path)
   if err != nil {
      return err
   }
   var rows justwatch.OfferRows
   for _, tag := range content.HrefLangTags {
      locale, ok := justwatch.EnUs.Locale(&tag)
      if !ok {
         return errors.New("Locale")
      }
      offers, err := tag.Offers(locale)
      if err != nil {
         return err
      }
      for _, offer := range offers {
         if offer.Monetization() {
            rows.Add(locale, offer)
         }
      }
      time.Sleep(f.sleep)
   }
   file, err := create(path.Base(url_path))
   if err != nil {
      return err
   }
   defer file.Close()
   slices.SortFunc(rows, func(a, b *justwatch.OfferRow) int {
      return strings.Compare(a.Url, b.Url)
   })
   _, err = fmt.Fprintln(file, rows)
   if err != nil {
      return err
   }
   return nil
}

func create(name string) (*os.File, error) {
   log.Println("Create", name)
   return os.Create(name)
}

type flag_set struct {
   address string
   sleep   time.Duration
}

func main() {
   http.DefaultClient.Transport = &justwatch.Transport
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
