package io

type reader interface {
   read() bool
}

func new_reader() reader {
   type hello struct{}
   return hello{}
}
