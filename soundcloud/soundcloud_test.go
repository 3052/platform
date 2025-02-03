package soundcloud

import (
   "fmt"
   "testing"
)

func TestTrack(t *testing.T) {
   fmt.Println(test_track)
}

var test_track = struct{
   id int64
   url string
}{
   id: 936653761,
   url: "https://soundcloud.com/kino-scmusic/mqymd53jtwag",
}

func TestNextData(t *testing.T) {
   var next next_data
   err := next.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", next)
}
