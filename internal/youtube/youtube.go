package main

import (
   "154.pages.dev/encoding"
   "154.pages.dev/log"
   "154.pages.dev/platform/youtube"
   "fmt"
   "log/slog"
   "net/http"
   "os"
   "slices"
)

func download(format *youtube.AdaptiveFormat, name string) error {
   ext, err := format.Ext()
   if err != nil {
      return err
   }
   file, err := os.Create(name + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   var meter log.ProgressMeter
   log.SetTransport(nil)
   ranges := format.Ranges()
   meter.Set(len(ranges))
   for _, byte_range := range ranges {
      err := func() error {
         res, err := http.Get(format.URL + byte_range)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         _, err = file.ReadFrom(meter.Reader(res))
         if err != nil {
            return err
         }
         return nil
      }()
      if err != nil {
         return err
      }
   }
   return nil
}

func (f flags) do_refresh() error {
   var code youtube.DeviceCode
   code.Post()
   fmt.Println(code)
   fmt.Scanln()
   token, err := code.Token()
   if err != nil {
      return err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   return os.WriteFile(home+"/youtube.json", token.Raw, 0666)
}
func (f flags) player() (*youtube.Player, error) {
   var token *youtube.Token
   switch f.request {
   case 0:
      f.r.Android()
   case 1:
      f.r.AndroidEmbed()
   case 2:
      f.r.AndroidCheck()
      home, err := os.UserHomeDir()
      if err != nil {
         return nil, err
      }
      token = new(youtube.Token)
      token.Raw, err = os.ReadFile(home + "/youtube.json")
      if err != nil {
         return nil, err
      }
      token.Unmarshal()
      if err := token.Refresh(); err != nil {
         return nil, err
      }
   }
   var play youtube.Player
   play.Post(f.r, token)
   return &play, nil
}

func (f flags) loop() error {
   play, err := f.player()
   if err != nil {
      return err
   }
   slog.Info("playability", "status", play.PlayabilityStatus)
   if f.itag >= 1 {
      var next youtube.WatchNext
      f.r.Web()
      if err := next.Post(f.r); err != nil {
         return err
      }
      format, ok := play.Format(f.itag)
      if ok {
         err := download(format, encoding.Name(next))
         if err != nil {
            return err
         }
      }
   } else {
      cmp := func(a, b youtube.AdaptiveFormat) int {
         return a.Bitrate - b.Bitrate
      }
      slices.SortFunc(play.StreamingData.AdaptiveFormats, cmp)
      for i, format := range play.StreamingData.AdaptiveFormats {
         if i >= 1 {
            fmt.Println()
         }
         fmt.Println(format)
      }
   }
   return nil
}

