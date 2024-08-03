package youtube

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

func TestFormat(t *testing.T) {
   var tube InnerTube
   tube.VideoId = android_ids[0]
   tube.Context.Client.ClientName = android
   play, err := tube.Player(nil)
   if err != nil {
      t.Fatal(err)
   }
   for _, format := range play.StreamingData.AdaptiveFormats {
      fmt.Println(format)
   }
}

func TestId(t *testing.T) {
   for _, test := range id_tests {
      var tube InnerTube
      err := tube.Set(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(tube.VideoId)
   }
}

var id_tests = []string{
   "https://youtube.com/shorts/9Vsdft81Q6w",
   "https://youtube.com/watch?v=XY-hOqcPGCY",
}

const image_test = "CUQ5ud4akt8"

func TestImage(t *testing.T) {
   for _, img := range Images {
      img.VideoId = image_test
      fmt.Println(img)
      resp, err := http.Head(img.String())
      if err != nil {
         t.Fatal(err)
      }
      if resp.StatusCode != http.StatusOK {
         t.Fatal(resp.Status)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
