package mp4

import (
   "encoding/hex"
   "io"
   "os"
   "testing"
)

func Test_Hulu(t *testing.T) {
   src, err := os.Open("enc.mp4")
   if err != nil {
      t.Fatal(err)
   }
   defer src.Close()
   key, err := hex.DecodeString("602a9289bfb9b1995b75ac63f123fc86")
   if err != nil {
      t.Fatal(err)
   }
   dec := make(Decrypt)
   dst, err := os.Create("dec.mp4")
   if err != nil {
      t.Fatal(err)
   }
   defer dst.Close()
   if err := dec.Init(io.LimitReader(src, 1530), dst); err != nil {
      t.Fatal(err)
   }
   if err := dec.Segment(src, dst, key); err != nil {
      t.Fatal(err)
   }
}
