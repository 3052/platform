package stream

import "time"

// youtube.com/watch?v=_ET7kEnDpyk
type video struct {
   season int64
   episode int64
   author string
   series string
   time time.Time
   title string
}
