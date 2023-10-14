package stream

import "time"

type Episode interface {
   Series() string
   Season() (int64, error)
   Episode() (int64, error)
   Title() string
}

type Film interface {
   Title() string
   Date() (time.Time, error)
}
