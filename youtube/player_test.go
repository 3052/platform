package youtube

import (
   "testing"
   "time"
)

func TestAndroid(t *testing.T) {
   for _, android_id := range android_ids {
      var tube InnerTube
      tube.Context.Client.Client.Name = "ANDROID"
      tube.VideoId = android_id
      play, err := tube.Player(nil)
      if err != nil {
         t.Fatal(err)
      }
      if play.PlayabilityStatus.Status != "OK" {
         t.Fatal(play)
      }
      if len(play.StreamingData.AdaptiveFormats) == 0 {
         t.Fatal("adaptiveFormats")
      }
      if play.VideoDetails.ViewCount == 0 {
         t.Fatal("viewCount")
      }
      time.Sleep(time.Second)
   }
}

var android_ids = []string{
   "H1BuwMTrtLQ", // content check
   "zv9NimPx3Es",
}

var embed_ids = []string{
   "HtVdAasjOgU",
   "WaOKSUlf4TM",
}

func TestAndroidEmbed(t *testing.T) {
   for _, embed_id := range embed_ids {
      var play Player
      tube := InnerTube{VideoId: embed_id}
      tube.AndroidEmbed()
      err := play.Post(tube, nil)
      if err != nil {
         t.Fatal(err)
      }
      if play.PlayabilityStatus.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
}

const web_id = "HPkDFc8hq5c"

func TestWeb(t *testing.T) {
   tube := InnerTube{VideoId: web_id}
   tube.Web()
   var p Player
   err := p.Post(tube, nil)
   if err != nil {
      t.Fatal(err)
   }
   if p.Author() == "" {
      t.Fatal("author")
   }
   if p.PlayabilityStatus.Reason != "" {
      t.Fatal("reason")
   }
   if p.PlayabilityStatus.Status != "OK" {
      t.Fatal("status")
   }
   if len(p.StreamingData.AdaptiveFormats) == 0 {
      t.Fatal("adaptiveFormats")
   }
   if _, err := p.Time(); err != nil {
      t.Fatal(err)
   }
   if p.Title() == "" {
      t.Fatal("title")
   }
   if p.VideoDetails.LengthSeconds <= 0 {
      t.Fatal("duration")
   }
   if p.VideoDetails.ShortDescription == "" {
      t.Fatal("shortDescription")
   }
   if p.VideoDetails.VideoId == "" {
      t.Fatal("videoId")
   }
   if p.VideoDetails.ViewCount <= 0 {
      t.Fatal("viewCount")
   }
}
