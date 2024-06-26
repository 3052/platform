package youtube

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

func TestModeLine(t *testing.T) {
   var req Request
   req.Android()
   req.VideoId = android_ids[0]
   var play Player
   if err := play.Post(req, nil); err != nil {
      t.Fatal(err)
   }
   for _, format := range play.StreamingData.AdaptiveFormats {
      fmt.Println(format)
   }
}

func TestId(t *testing.T) {
   for _, test := range id_tests {
      var req Request
      err := req.Set(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(req.VideoId)
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
