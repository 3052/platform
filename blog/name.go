package stream

import (
   "strconv"
   "time"
)

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

type Video interface {
   Author() string
   Title() string
}

func Format_Episode(e Episode) (string, error) {
   b := []byte(e.Series())
   b = append(b, " - S"...)
   if v, err := e.Season(); err != nil {
      return "", err
   } else {
      b = strconv.AppendInt(b, v, 10)
   }
   b = append(b, " E"...)
   if v, err := e.Episode(); err != nil {
      return "", err
   } else {
      b = strconv.AppendInt(b, v, 10)
   }
   b = append(b, " - "...)
   b = append(b, e.Title()...)
   clean(b)
   return string(b), nil
}

func Format_Film(f Film) (string, error) {
   b := []byte(f.Title())
   b = append(b, " - "...)
   if v, err := f.Date(); err != nil {
      return "", err
   } else {
      b = v.AppendFormat(b, "2006")
   }
   clean(b)
   return string(b), nil
}

func clean(path []byte) {
   m := map[byte]bool{
      '"': true,
      '*': true,
      '/': true,
      ':': true,
      '<': true,
      '>': true,
      '?': true,
      '\\': true,
      '|': true,
   }
   for k, v := range path {
      if m[v] {
         path[k] = '-'
      }
   }
}
