package main

import (
   "154.pages.dev/platform/youtube"
   "154.pages.dev/text"
   "fmt"
   "log/slog"
   "net/http"
   "os"
   "slices"
)

func (f flags) do_refresh() error {
   var code youtube.DeviceCode
   code.New()
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
   return os.WriteFile(home+"/youtube.json", token.Data, 0666)
}

func (f flags) player() (*youtube.Player, error) {
   var token *youtube.AuthToken
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
      token = new(youtube.AuthToken)
      token.Data, err = os.ReadFile(home + "/youtube.json")
      if err != nil {
         return nil, err
      }
      err = token.Unmarshal()
      if err != nil {
         return nil, err
      }
      err = token.Refresh()
      if err != nil {
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
   slog.Info("playability", "status", play.PlayabilityStatus.Status)
   // download one
   for _, format := range play.StreamingData.AdaptiveFormats {
      if format.Itag == f.itag {
         return f.download(format, play.VideoDetails.Title)
      }
   }
   // print all
   slices.SortFunc(
      play.StreamingData.AdaptiveFormats, youtube.AdaptiveFormat.CompareBitrate,
   )
   for i, format := range play.StreamingData.AdaptiveFormats {
      if i >= 1 {
         fmt.Println()
      }
      fmt.Println(format)
   }
   return nil
}

func (f flags) download(format youtube.AdaptiveFormat, name string) error {
   ext, err := format.Ext()
   if err != nil {
      return err
   }
   file, err := os.Create(text.Clean(name) + ext)
   if err != nil {
      return err
   }
   defer file.Close()
   var meter text.ProgressMeter
   ranges := format.Ranges()
   meter.Set(len(ranges))
   for _, byte_range := range ranges {
      err := func() error {
         resp, err := http.Get(format.Url + byte_range)
         if err != nil {
            return err
         }
         defer resp.Body.Close()
         _, err = file.ReadFrom(meter.Reader(resp))
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
