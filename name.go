package stream

import (
   "strconv"
   "time"
)

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

type Episode_Name interface {
   Series() string
   Season() (int64, error)
   Episode() (int64, error)
   Title() string
}

func Episode(e Episode_Name) (string, error) {
   var b []byte
   b = append(b, e.Series()...)
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

type Film_Name interface {
   Title() string
   Date() (time.Time, error)
}

func Film(f Film_Name) (string, error) {
   var b []byte
   b = append(b, f.Title()...)
   b = append(b, " - "...)
   if v, err := f.Date(); err != nil {
      return "", err
   } else {
      b = v.AppendFormat(b, "2006")
   }
   clean(b)
   return string(b), nil
}

type Video_Name interface {
   Author() string
   Title() string
}

// wikipedia.org/wiki/Charlie_Bit_My_Finger
func Video(v Video_Name) string {
   var b []byte
   b = append(b, v.Author()...)
   b = append(b, " - "...)
   b = append(b, v.Title()...)
   clean(b)
   return string(b)
}
