package stream

type Namer interface {
   Owner() (string, bool)
   Show() (string, bool)
   Season() (string, bool)
   Episode() (string, bool)
   Title() (string, bool)
   Year() (string, bool)
}

// Owner Show Season Episode - Title - Year
func Name(n Namer) string {
   var b []byte
   if v, ok := n.Owner(); ok {
      b = append(b, v...)
   }
   if v, ok := n.Show(); ok {
      if b != nil {
         b = append(b, ' ')
      }
      b = append(b, v...)
   }
   if v, ok := n.Season(); ok {
      if b != nil {
         b = append(b, ' ')
      }
      b = append(b, v...)
   }
   if v, ok := n.Episode(); ok {
      if b != nil {
         b = append(b, ' ')
      }
      b = append(b, v...)
   }
   if v, ok := n.Title(); ok {
      if b != nil {
         b = append(b, " - "...)
      }
      b = append(b, v...)
   }
   if v, ok := n.Year(); ok {
      if b != nil {
         b = append(b, " - "...)
      }
      b = append(b, v...)
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
