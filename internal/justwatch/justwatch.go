package main

import (
   "154.pages.dev/platform/justwatch"
   "fmt"
   "slices"
   "time"
)

func (f flags) stream() error {
   content, err := f.path.Content()
   if err != nil {
      return err
   }
   var groups justwatch.OfferGroups
   for _, tag := range content.HrefLangTags {
      locale, ok := justwatch.EnglishLocales.Locale(tag)
      if !ok {
         return justwatch.LocaleState{}
      }
      offers, err := tag.Offers(locale)
      if err != nil {
         return err
      }
      if !f.all {
         offers = slices.DeleteFunc(offers, justwatch.Url)
      }
      for _, offer := range slices.DeleteFunc(offers, justwatch.Monetization) {
         groups.Add(locale, offer)
      }
      time.Sleep(f.sleep)
   }
   fmt.Println(groups)
   return nil
}
