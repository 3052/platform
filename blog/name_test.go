package stream

import (
   "fmt"
   "testing"
   "time"
)

type tester struct{}

func (tester) Series() string {
   return "series"
}

func (tester) Season() (int64, error) {
   return 2, nil
}

func (tester) Episode() (int64, error) {
   return 3, nil
}

func (tester) Title() string {
   return "title"
}

func (tester) Date() (time.Time, error) {
   return time.Now(), nil
}

func Test_Name(t *testing.T) {
   var test tester
   s, err := func() (string, error) {
      if true {
         return Format_Episode(test)
      }
      return Format_Film(test)
   }()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(s)
}

