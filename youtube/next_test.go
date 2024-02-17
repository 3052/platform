package youtube

import (
   "154.pages.dev/rosso"
   "fmt"
   "testing"
   "time"
)

// youtube.com/channel/UCuVPpxrm2VAgpH3Ktln4HXg
var video_ids = []string{
   "7KLCti7tOXE", // video
   "2ZcDwdXEVyI", // episode
   "PBcnZCa1dEk", // film
}

func TestNext(t *testing.T) {
   for _, video_id := range video_ids {
      req := Request{VideoId: video_id}
      req.Web()
      var next WatchNext
      err := next.Post(req)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(rosso.Name(next))
      time.Sleep(99*time.Millisecond)
   }
}
