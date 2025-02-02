package main

import (
   "41.neocities.org/platform/justwatch"
   "bytes"
   "errors"
   "flag"
   "fmt"
   "io"
   "log"
   "net/http"
   "slices"
   "time"
)

func main() {
   http.DefaultClient.Transport = transport{}
   log.SetFlags(log.Ltime)
   var f flags
   flag.Var(&f.address, "a", "address")
   flag.DurationVar(&f.sleep, "s", 99*time.Millisecond, "sleep")
   flag.BoolVar(&f.filter, "f", false, "filter")
   flag.Parse()
   if f.address.String() != "" {
      err := f.stream()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

func (transport) RoundTrip(req *http.Request) (*http.Response, error) {
   if req.Body != nil {
      data, err := io.ReadAll(req.Body)
      if err != nil {
         return nil, err
      }
      req.Body.Close()
      req.Body = io.NopCloser(bytes.NewReader(data))
      log.Print(string(data))
   } else {
      log.Print(req.URL)
   }
   return http.DefaultTransport.RoundTrip(req)
}

type transport struct{}

type flags struct {
   address justwatch.Address
   filter  bool
   sleep   time.Duration
}

func (f *flags) stream() error {
   content, err := f.address.Content()
   if err != nil {
      return err
   }
   var groups justwatch.OfferGroups
   for _, tag := range content.HrefLangTags {
      locale, ok := justwatch.EnglishLocales.Locale(&tag)
      if !ok {
         return errors.New("EnglishLocales.Locale")
      }
      offers, err := tag.Offers(locale)
      if err != nil {
         return err
      }
      if f.filter {
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
