package bandcamp

import (
   "fmt"
   "testing"
)

const art_id = 3809045440

func Test_Image(t *testing.T) {
   for _, img := range Images {
      ref := img.URL(art_id)
      fmt.Println(ref)
   }
}
