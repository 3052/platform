package main

import (
   "154.pages.dev/platform/justwatch"
   "154.pages.dev/text"
   "errors"
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
   text.SetTransport(nil)
   for _, tag := range content.HrefLangTags {
      fmt.Println(tag.Href)
      locale, ok := justwatch.EnglishLocales.Locale(tag)
      if !ok {
         return errors.New("justwatch.LocaleStates.Locale")
      }
      offers, err := tag.Offers(locale)
      if err != nil {
         return err
      }
      if f.filter {
         offers = slices.DeleteFunc(offers, justwatch.URL)
      }
      for _, offer := range slices.DeleteFunc(offers, justwatch.Monetization) {
         groups.Add(locale, offer)
      }
      time.Sleep(f.sleep)
   }
   fmt.Println(groups)
   return nil
}
