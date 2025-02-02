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

func (f *flags) stream() error {
   content, err := f.address.Content()
   if err != nil {
      return err
   }
   var rows justwatch.OfferRows
   for _, tag := range content.HrefLangTags {
      locale, ok := justwatch.English.Locale(&tag)
      if !ok {
         return errors.New("Locale")
      }
      offers, err := tag.Offers(locale)
      if err != nil {
         return err
      }
      offers = slices.DeleteFunc(offers, justwatch.Offer.Monetization)
      for _, offer := range offers {
         rows.Add(locale, &offer)
      }
      time.Sleep(f.sleep)
   }
   fmt.Println(rows)
   return nil
}

func main() {
   http.DefaultClient.Transport = transport{}
   log.SetFlags(log.Ltime)
   var f flags
   flag.Var(&f.address, "a", "address")
   flag.DurationVar(&f.sleep, "s", 99*time.Millisecond, "sleep")
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
   sleep   time.Duration
}
