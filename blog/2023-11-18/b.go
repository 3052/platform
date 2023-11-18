package io

type reader interface {
   read() bool
}

func new_reader() reader {
   type hello struct{}
   func (hello) read() bool { return true }
   return hello{}
}
