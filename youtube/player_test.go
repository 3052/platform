package youtube

import (
   "testing"
   "time"
)

func TestWeb(t *testing.T) {
   var tube InnerTube
   tube.VideoId = web_id
   tube.Context.Client.ClientName = web
   play, err := tube.Player(nil)
   if err != nil {
      t.Fatal(err)
   }
   if play.Microformat.PlayerMicroformatRenderer.PublishDate.Time.IsZero() {
      t.Fatal("PublishDate")
   }
   if play.PlayabilityStatus.Reason != "" {
      t.Fatal("reason")
   }
   if play.PlayabilityStatus.Status != "OK" {
      t.Fatal("status")
   }
   if len(play.StreamingData.AdaptiveFormats) == 0 {
      t.Fatal("adaptiveFormats")
   }
   if play.VideoDetails.Author == "" {
      t.Fatal("author")
   }
   if play.VideoDetails.LengthSeconds <= 0 {
      t.Fatal("duration")
   }
   if play.VideoDetails.ShortDescription == "" {
      t.Fatal("shortDescription")
   }
   if play.VideoDetails.Title == "" {
      t.Fatal("title")
   }
   if play.VideoDetails.VideoId == "" {
      t.Fatal("videoId")
   }
   if play.VideoDetails.ViewCount <= 0 {
      t.Fatal("viewCount")
   }
}

func TestAndroid(t *testing.T) {
   for _, android_id := range android_ids {
      var tube InnerTube
      tube.Context.Client.ClientName = android
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
      var tube InnerTube
      tube.VideoId = embed_id
      tube.Context.Client.ClientName = android_embedded_player
      play, err := tube.Player(nil)
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
