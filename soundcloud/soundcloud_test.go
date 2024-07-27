package soundcloud

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
)
package soundcloud

import (
   "encoding/json"
   "os"
   "testing"
)

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
const (
   id = 936653761
   address = "https://soundcloud.com/kino-scmusic/mqymd53jtwag"
)

func TestResolve(t *testing.T) {
   track, err := Resolve(address)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(track)
}

func TestTrack(t *testing.T) {
   var track ClientTrack
   err := track.New(id)
   if err != nil {
      t.Fatal(err)
   }
   media, err := track.Progressive()
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   enc.Encode(media)
}
