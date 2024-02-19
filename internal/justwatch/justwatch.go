package main

import (
   "154.pages.dev/log"
   "blog/justwatch"
   "flag"
   "log/slog"
   "net/url"
   "time"
)

type flags struct {
   address string
   level log.Level
   sleep time.Duration
}

func main() {
   var f flags
   flag.StringVar(&f.address, "a", "", "address")
   flag.DurationVar(&f.sleep, "s", 99*time.Millisecond, "sleep")
   flag.TextVar(&f.level, "v", f.level, "log level")
   flag.Parse()
   log.Set_Logger(f.level)
   if f.address != "" {
      err := f.stream()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

func (f flags) stream() error {
   address, err := url.Parse(f.address)
   if err != nil {
      return err
   }
   content, err := justwatch.New_URLs(address.Path)
   if err != nil {
      return err
   }
   offer := make(justwatch.Offers)
   for _, tag := range content.Href_Lang_Tags {
      slog.Debug(tag.Href_Lang)
      if !justwatch.Blacklist[tag.Href_Lang] {
         v := tag.Variables()
         text, err := v.Text()
         if err != nil {
            return err
         }
         slog.Info(text)
         detail, err := v.Details()
         if err != nil {
            return err
         }
         offer.Add(v.Country_Code, detail)
         time.Sleep(f.sleep)
      }
   }
   text, err := offer.Stream().Text()
   if err != nil {
      return err
   }
   slog.Info(text)
   return nil
}
