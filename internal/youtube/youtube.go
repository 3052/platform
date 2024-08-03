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

func (f *flags) loop() error {
   var token *youtube.AuthToken
   switch f.request {
   case 0:
      f.r.Android()
   case 1:
      f.r.AndroidEmbed()
   case 2:
      text, err := os.ReadFile(f.home + "/youtube.json")
      if err != nil {
         return err
      }
      token = new(youtube.AuthToken)
      err = token.Unmarshal(text)
      if err != nil {
         return err
      }
      err = token.Refresh()
      if err != nil {
         return err
      }
      f.r.AndroidCheck()
   }
   var play youtube.Player
   err := play.New(f.r, token)
   if err != nil {
      return err
   }
   slog.Info("playability", "status", play.PlayabilityStatus.Status)
   // download one
   for _, format := range play.StreamingData.AdaptiveFormats {
      if format.Itag == f.itag {
         return download(format, play.VideoDetails.Title)
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

func download(format youtube.AdaptiveFormat, name string) error {
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
func (f *flags) New() error {
   var err error
   f.home, err = os.UserHomeDir()
   if err != nil {
      return err
   }
   return nil
}

func write_code() error {
   var code youtube.DeviceCode
   err := code.New()
   if err != nil {
      return err
   }
   text, err := code.Marshal()
   if err != nil {
      return err
   }
   fmt.Println(code)
   return os.WriteFile("code.json", text, 0666)
}

func (f *flags) write_token() error {
   text, err := os.ReadFile("code.json")
   if err != nil {
      return err
   }
   var code youtube.DeviceCode
   err = code.Unmarshal(text)
   if err != nil {
      return err
   }
   token, err := code.Token()
   if err != nil {
      return err
   }
   text, err = token.Marshal()
   if err != nil {
      return err
   }
   return os.WriteFile(f.home + "/youtube.json", text, 0666)
}
