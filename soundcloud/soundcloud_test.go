package soundcloud

import (
   "encoding/json"
   "fmt"
   "net/http"
   "os"
   "strings"
   "testing"
   "time"
)

func TestTrack(t *testing.T) {
   var track ClientTrack
   err := track.New(test_track.id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", track)
}

var test_track = struct{
   id int64
   url string
}{
   id: 936653761,
   url: "https://soundcloud.com/kino-scmusic/mqymd53jtwag",
}

func TestResolve(t *testing.T) {
   track, err := Resolve(test_track.url)
   if err != nil {
      t.Fatal(err)
   }
   media, err := track.Progressive()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(media)
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
      time.Sleep(time.Second)
   }
}

func TestNextData(t *testing.T) {
   var next next_data
   err := next.New()
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(next)
}
