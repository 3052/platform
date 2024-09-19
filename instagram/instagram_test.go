package instagram

import (
   "fmt"
   "testing"
   "time"
)

var tests = []string{
   "instagram.com/jadelynmusic/p/C_1CZn2y6Lt",
   "instagram.com/jadelynmusic/p/DABx8RVSMuO",
}

func TestMedia(t *testing.T) {
   for i, test := range tests {
      if i >= 1 {
         fmt.Println()
      }
      var web Address
      web.Set(test)
      media, err := web.Media()
      if err != nil {
         t.Fatal(err)
      }
      for _, display := range media.DisplayUrls() {
         fmt.Println(display)
      }
      time.Sleep(time.Second)
   }
}
