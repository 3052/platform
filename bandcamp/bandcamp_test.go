package bandcamp

import (
   "fmt"
   "reflect"
   "testing"
   "time"
)

func TestSize(t *testing.T) {
   size := reflect.TypeOf(&struct{}{}).Size()
   for _, test := range size_tests {
      if reflect.TypeOf(test).Size() > size {
         fmt.Printf("*%T\n", test)
      } else {
         fmt.Printf("%T\n", test)
      }
   }
}

var size_tests = []any{
   AlbumTrack{},
   BandDetails{},
   Image{},
   Item{},
   ReportParams{},
   Tralbum{},
   invalid_type{},
}

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
