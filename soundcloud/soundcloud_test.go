package soundcloud

import (
   "fmt"
   "net/http"
   "strings"
   "testing"
   "time"
)

func TestResolve(t *testing.T) {
   var track ClientTrack
   err := track.Resolve(test_track.url)
   if err != nil {
      t.Fatal(err)
   }
   progressive, ok := track.Progressive()
   if !ok {
      t.Fatal("ClientTrack.Progressive")
   }
   media, err := progressive.Media()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", media)
}

func TestNextData(t *testing.T) {
   var next next_data
   err := next.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", next)
}

var test_track = struct{
   id int64
   url string
}{
   id: 936653761,
   url: "https://soundcloud.com/kino-scmusic/mqymd53jtwag",
}

func TestArtwork(t *testing.T) {
   for _, artwork := range artworks {
      address := func() string {
         var b strings.Builder
         b.WriteString("https://i1.sndcdn.com/artworks-000365245539-x9ixki-")
         b.WriteString(artwork.size)
         b.WriteString(".jpg")
         return b.String()
      }()
      resp, err := http.Head(address)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(resp.Status)
      time.Sleep(99 * time.Millisecond)
   }
}

func TestTrack(t *testing.T) {
   var track ClientTrack
   err := track.New(test_track.id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", track)
}
