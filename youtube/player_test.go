package youtube

import (
   "testing"
   "time"
)

var android_ids = []string{
   "H1BuwMTrtLQ", // content check
   "zv9NimPx3Es",
}

func TestAndroid(t *testing.T) {
   for _, android_id := range android_ids {
      var p Player
      req := Request{VideoId: android_id}
      req.Android()
      err := p.Post(req, nil)
      if err != nil {
         t.Fatal(err)
      }
      if p.PlayabilityStatus.Status != "OK" {
         t.Fatal(p)
      }
      if len(p.StreamingData.AdaptiveFormats) == 0 {
         t.Fatal("adaptiveFormats")
      }
      if p.VideoDetails.ViewCount == 0 {
         t.Fatal("viewCount")
      }
      time.Sleep(time.Second)
   }
}

var embed_ids = []string{
   "HtVdAasjOgU",
   "WaOKSUlf4TM",
}

func TestAndroidEmbed(t *testing.T) {
   for _, embed_id := range embed_ids {
      var play Player
      req := Request{VideoId: embed_id}
      req.AndroidEmbed()
      err := play.Post(req, nil)
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
   r := Request{VideoId: web_id}
   r.Web()
   var p Player
   err := p.Post(r, nil)
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
