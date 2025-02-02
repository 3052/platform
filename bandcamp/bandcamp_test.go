package bandcamp

import (
   "fmt"
   "testing"
)

func TestUrl(t *testing.T) {
   fmt.Println(tests)
}

var tests = []string{
   "https://schnaussandmunk.bandcamp.com",
   "https://schnaussandmunk.bandcamp.com/album/passage-2",
   "https://schnaussandmunk.bandcamp.com/music",
   "https://schnaussandmunk.bandcamp.com/track/amaris-2",
}
