package main

import (
   "154.pages.dev/platform/justwatch"
   "154.pages.dev/text"
   "flag"
   "fmt"
   "slices"
   "time"
)

func (f *flags) stream() error {
   content, err := f.address.Content()
   if err != nil {
      return err
   }
   var groups justwatch.OfferGroups
   for _, tag := range content.HrefLangTags {
      locale, ok := justwatch.EnglishLocales.Locale(&tag)
      if !ok {
         return justwatch.EnglishLocales.LocaleError()
      }
      offers, err := tag.Offers(locale)
      if err != nil {
         return err
      }
      if !f.all {
         offers = slices.DeleteFunc(offers, justwatch.Url)
      }
      for _, offer := range slices.DeleteFunc(offers, justwatch.Monetization) {
         groups.Add(&offer, locale)
      }
      time.Sleep(f.sleep)
   }
   fmt.Println(groups)
   return nil
}

type flags struct {
   all bool
   sleep time.Duration
   address justwatch.Address
}

func main() {
   var f flags
   flag.Var(&f.address, "a", "address")
   flag.DurationVar(&f.sleep, "s", 99*time.Millisecond, "sleep")
   flag.BoolVar(&f.all, "all", false, "all results")
   flag.Parse()
   text.Transport{}.Set(true)
   if f.address.Path != "" {
      err := f.stream()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
