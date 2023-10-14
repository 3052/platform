package stream

import (
   "fmt"
   "time"
   "testing"
)

func Test_Name(t *testing.T) {
   //var test tester
   var _ tester
   name, err := func() (string, error) {
      /*
      if Film(test) {
         return Film(test)
      }
      return Episode(test)
      */
      return "", nil
   }()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(name)
}

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
