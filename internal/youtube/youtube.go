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
   ranges := format.Ranges()
   var meter text.ProgressMeter
   meter.Set(len(ranges))
   text.Transport{}.Set(false)
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

func code() error {
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
   return os.WriteFile("code.txt", text, 0666)
}

func (f *flags) token() error {
   text, err := os.ReadFile("code.txt")
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
   return os.WriteFile(f.home + "/youtube.txt", text, 0666)
}

func (f *flags) loop() error {
   var token *youtube.AuthToken
   if f.read_token {
      text, err := os.ReadFile(f.home + "/youtube.txt")
      if err != nil {
         return err
      }
      token = &youtube.AuthToken{}
      err = token.Unmarshal(text)
      if err != nil {
         return err
      }
      err = token.Refresh()
      if err != nil {
         return err
      }
   }
   play, err := f.tube.Player(token)
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
      fmt.Println(&format)
   }
   return nil
}
