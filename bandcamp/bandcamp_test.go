package bandcamp

import (
   "fmt"
   "testing"
   "time"
)

var tests = []string{
   "https://schnaussandmunk.bandcamp.com",
   "https://schnaussandmunk.bandcamp.com/album/passage-2",
   "https://schnaussandmunk.bandcamp.com/music",
   "https://schnaussandmunk.bandcamp.com/track/amaris-2",
}

func TestParam(t *testing.T) {
   for _, test := range tests {
      var params ReportParams
      err := params.New(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", params)
      time.Sleep(99 * time.Millisecond)
   }
}
const art_id = 3809045440

func TestImage(t *testing.T) {
   for _, img := range Images {
      ref := img.URL(art_id)
      fmt.Println(ref)
   }
}
