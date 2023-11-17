package dash

import (
   "encoding/base64"
   "encoding/hex"
   "encoding/xml"
   "io"
   "strconv"
   "strings"
)

type Representation struct {
   Adaptation *Adaptation
   Bandwidth int `xml:"bandwidth,attr"`
   Base_URL string `xml:"BaseURL"`
   Codecs string `xml:"codecs,attr"`
   Content_Protection []Content_Protection `xml:"ContentProtection"`
   Height int `xml:"height,attr"`
   ID string `xml:"id,attr"`
   MIME_Type string `xml:"mimeType,attr"`
   // zero-indexed & inclusive
   // developer.mozilla.org/docs/Web/HTTP/Headers/Range
   Segment_Base struct {
      Index_Range string `xml:"indexRange,attr"`
   } `xml:"SegmentBase"`
   Segment_Template *Segment_Template `xml:"SegmentTemplate"`
   Width int `xml:"width,attr"`
} 

func (r Representation) KID_PSSH() ([]byte, []byte, error) {
   var kid []byte
   var pssh []byte
   for _, c := range r.Content_Protection {
      var err error
      switch c.Scheme_ID_URI {
      case "urn:mpeg:dash:mp4protection:2011":
         c.Default_KID = strings.ReplaceAll(c.Default_KID, "-", "")
         kid, err = hex.DecodeString(c.Default_KID)
         if err != nil {
            return nil, nil, err
         }
      case "urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed":
         pssh, err = base64.StdEncoding.DecodeString(c.PSSH)
         if err != nil {
            return nil, nil, err
         }
      }
   }
   return kid, pssh, nil
}

type Content_Protection struct {
   Default_KID string `xml:"default_KID,attr"`
   PSSH string `xml:"pssh"`
   Scheme_ID_URI string `xml:"schemeIdUri,attr"`
} 

func (r Representation) Initialization() (string, bool) {
   if s := r.Segment_Template; s != nil {
      if i := s.Initialization; i != "" {
         return strings.Replace(i, "$RepresentationID$", r.ID, 1), true
      }
   }
   return "", false
}

func (r Representation) Ext() (string, bool) {
   switch {
   case Audio(r):
      return ".m4a", true
   case Video(r):
      return ".m4v", true
   }
   return "", false
}

func Not[E any](fn func(E) bool) func(E) bool {
   return func(value E) bool {
      return !fn(value)
   }
}

func (r Representation) Media() []string {
   if r.Segment_Template == nil {
      return nil
   }
   replace := func(s string, i int) string {
      s = strings.Replace(s, "$Number$", strconv.Itoa(i), 1)
      return strings.Replace(s, "$RepresentationID$", r.ID, 1)
   }
   var media []string
   for _, segment := range r.Segment_Template.Segment_Timeline.S {
      segment.T = r.Segment_Template.Start_Number
      for segment.R >= 0 {
         medium := replace(r.Segment_Template.Media, segment.T)
         media = append(media, medium)
         r.Segment_Template.Start_Number++
         segment.R--
         segment.T++
      }
   }
   return media
}

type Segment_Template struct {
   Initialization string `xml:"initialization,attr"`
   Media string `xml:"media,attr"`
   Segment_Timeline struct {
      S []struct {
         D int `xml:"d,attr"` // duration
         R int `xml:"r,attr"` // repeat
         T int `xml:"t,attr"` // time
      }
   } `xml:"SegmentTimeline"`
   Start_Number int `xml:"startNumber,attr"`
}

type Adaptation struct {
   Content_Protection []Content_Protection `xml:"ContentProtection"`
   Lang string `xml:"lang,attr"`
   MIME_Type string `xml:"mimeType,attr"`
   Segment_Template *Segment_Template `xml:"SegmentTemplate"`
   Representation []Representation
   Role *struct {
      Value string `xml:"value,attr"`
   }
}

func Audio(r Representation) bool {
   return r.MIME_Type == "audio/mp4"
}

func Video(r Representation) bool {
   return r.MIME_Type == "video/mp4"
}

func (r Representation) String() string {
   var s []string
   if r.Width >= 1 {
      s = append(s, "width: " + strconv.Itoa(r.Width))
   }
   if r.Height >= 1 {
      s = append(s, "height: " + strconv.Itoa(r.Height))
   }
   if r.Bandwidth >= 1 {
      s = append(s, "bandwidth: " + strconv.Itoa(r.Bandwidth))
   }
   if r.Codecs != "" {
      s = append(s, "codecs: " + r.Codecs)
   }
   s = append(s, "type: " + r.MIME_Type)
   if r.Adaptation.Role != nil {
      s = append(s, "role: " + r.Adaptation.Role.Value)
   }
   if r.Adaptation.Lang != "" {
      s = append(s, "language: " + r.Adaptation.Lang)
   }
   return strings.Join(s, "\n")
}

func Representations(r io.Reader) ([]Representation, error) {
   var s struct {
      Period struct {
         Adaptation_Set []Adaptation `xml:"AdaptationSet"`
      }
   }
   err := xml.NewDecoder(r).Decode(&s)
   if err != nil {
      return nil, err
   }
   var reps []Representation
   for _, ada := range s.Period.Adaptation_Set {
      ada := ada
      for _, rep := range ada.Representation {
         rep := rep
         rep.Adaptation = &ada
         if rep.Content_Protection == nil {
            rep.Content_Protection = ada.Content_Protection
         }
         if rep.Segment_Template == nil {
            rep.Segment_Template = ada.Segment_Template
         }
         if rep.MIME_Type == "" {
            rep.MIME_Type = ada.MIME_Type
         }
         reps = append(reps, rep)
      }
   }
   return reps, nil
}
