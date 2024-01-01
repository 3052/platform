package stream

type Namer interface {
   Episode() (string, bool)
   Owner() (string, bool)
   Release_Date() (string, bool)
   Season() (string, bool)
   Show() (string, bool)
   Title() (string, bool)
}

func Name(n Namer) string {
   var b []byte
   date, date_ok := n.Release_Date()
   show, show_ok := n.Show()
   if !date_ok {
      if v, ok := n.Owner(); ok {
         b = append(b, v...)
      }
   }
   if show_ok {
      b = append(b, show...)
   }
   if v, ok := n.Season(); ok {
      b = append(b, ' ')
      b = append(b, v...)
   }
   if v, ok := n.Episode(); ok {
      b = append(b, ' ')
      b = append(b, v...)
   }
   if v, ok := n.Title(); ok {
      if b != nil {
         b = append(b, " - "...)
      }
      b = append(b, v...)
   }
   if !show_ok {
      if date_ok {
         b = append(b, " - "...)
         b = append(b, date...)
      }
   }
   clean(b)
   return string(b)
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
