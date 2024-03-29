package main

import (
   "154.pages.dev/log"
   "154.pages.dev/platform/justwatch"
   "errors"
   "flag"
   "fmt"
   "log/slog"
   "net/url"
   "slices"
   "time"
)

func (f flags) stream() error {
   address, err := url.Parse(f.address)
   if err != nil {
      return err
   }
   var content justwatch.ContentUrls
   if err := content.New(address.Path); err != nil {
      return err
   }
   var groups justwatch.OfferGroups
   log.SetTransport(nil)
   for _, tag := range content.Href_Lang_Tags {
      slog.Info(tag.Href)
      locale, ok := justwatch.EnglishLocales.Locale(tag)
      if !ok {
         return errors.New("justwatch.LocaleStates.Locale")
      }
      offers, err := tag.Offers(locale)
      if err != nil {
         return err
      }
      for _, offer := range slices.DeleteFunc(offers, justwatch.Delete) {
         groups.Add(locale, offer)
      }
      time.Sleep(f.sleep)
   }
   fmt.Println(groups)
   return nil
}

type flags struct {
   address string
   sleep time.Duration
   v log.Level
}

func main() {
   var f flags
   flag.StringVar(&f.address, "a", "", "address")
   flag.DurationVar(&f.sleep, "s", 99*time.Millisecond, "sleep")
   flag.TextVar(&f.v.Level, "v", f.v.Level, "log level")
   flag.Parse()
   f.v.Set()
   log.Transport{}.Set()
   if f.address != "" {
      err := f.stream()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

