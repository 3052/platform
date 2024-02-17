package main

import (
   "154.pages.dev/log"
   "154.pages.dev/media/youtube"
   "154.pages.dev/rosso"
   "fmt"
   "log/slog"
   "net/http"
   "os"
   "text/template"
)

func (f flags) loop() error {
   play, err := f.player()
   if err != nil {
      return err
   }
   slog.Info("playability", "status", play.PlayabilityStatus)
   formats := play.StreamingData.AdaptiveFormats
   if len(f.itag) == 0 {
      html, err := new(template.Template).Parse(youtube.ModeLine)
      if err != nil {
         return err
      }
      return html.Execute(os.Stdout, formats)
   }
   var next youtube.WatchNext
   f.r.Web()
   if err := next.Post(f.r); err != nil {
      return err
   }
   for _, format := range formats {
      if _, ok := f.itag[format.Itag]; ok {
         err := download(format, rosso.Name(next))
         if err != nil {
            return err
         }
      }
   }
   return nil
}

func download(format youtube.Format, name string) error {
   ext, err := format.Ext()
   if err != nil {
      return err
   }
   file, err := os.Create(name + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   log.TransportDebug()
   ranges := format.Ranges()
   var meter log.ProgressMeter
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
      var token youtube.Token
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
