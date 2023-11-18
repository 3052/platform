package main

import "fmt"

type Representation int

func main() {
   var a Representation
   fmt.Printf("%#v\n", a)
   type SegmentTemplate struct {
      StartNumber int `xml:"startNumber,attr"`
   }
   type AdaptationSet struct {
      Representation []Representation
      SegmentTemplate *SegmentTemplate
   }
   type Representation struct {
      AdaptationSet *AdaptationSet
      SegmentTemplate *SegmentTemplate
   }
   var b Representation
   fmt.Printf("%#v\n", b)
}
