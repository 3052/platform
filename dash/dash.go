package dash

import (
   "encoding/base64"
   "encoding/hex"
   "encoding/xml"
   "io"
   "strconv"
   "strings"
)

func (a Adaptation) Type() string {
   if a.MIME_Type != "" {
      return a.MIME_Type
   }
   return a.Content_Type
}

func (a Adaptation) String() string {
   var s []string
   for i, r := range a.Representation {
      if i >= 1 {
         s = append(s, "")
      }
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
      s = append(s, "type: " + a.Type())
      if a.Role != nil {
         s = append(s, "role: " + a.Role.Value)
      }
      if a.Lang != "" {
         s = append(s, "language: " + a.Lang)
      }
   }
   return strings.Join(s, "\n")
}

func (s Segment_Base) Start() (int64, error) {
   i := strings.Index(s.Index_Range, "-")
   return strconv.ParseInt(s.Index_Range[:i], 10, 64)
}

type Segment_Base struct {
   Index_Range string `xml:"indexRange,attr"`
}

type Content_Protection struct {
   Default_KID string `xml:"default_KID,attr"`
   PSSH string `xml:"pssh"`
   Scheme_ID_URI string `xml:"schemeIdUri,attr"`
} 

type Segment_Template struct {
   Initialization string `xml:"initialization,attr"`
   Media string `xml:"media,attr"`
   Start_Number int `xml:"startNumber,attr"`
   Segment_Timeline struct {
      S []struct {
         D int `xml:"d,attr"` // duration
         R int `xml:"r,attr"` // repeat
         T int `xml:"t,attr"` // time
      }
   } `xml:"SegmentTimeline"`
}

func (a Adaptation) Audio() bool {
   return strings.HasPrefix(a.Type(), "audio")
}

func (a Adaptation) Video() bool {
   return strings.HasPrefix(a.Type(), "video")
}

func (a Adaptation) Ext() (string, bool) {
   switch {
   case a.Audio():
      return ".m4a", true
   case a.Video():
      return ".m4v", true
   }
   return "", false
}

func (r Representation) Initialization() (string, bool) {
   if s := r.Segment_Template; s != nil {
      if i := s.Initialization; i != "" {
         return strings.Replace(i, "$RepresentationID$", r.ID, 1), true
      }
   }
   return "", false
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
      for _, rep := range ada.Representation {
         if rep.Content_Protection == nil {
            rep.Content_Protection = ada.Content_Protection
         }
         if rep.Segment_Template == nil {
            rep.Segment_Template = ada.Segment_Template
         }
         reps = append(reps, rep)
      }
   }
   return reps, nil
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

type Adaptation struct {
   Lang string `xml:"lang,attr"`
   Role *struct {
      Value string `xml:"value,attr"`
   }
   Representation []Representation
   Content_Type string `xml:"contentType,attr"`
   MIME_Type string `xml:"mimeType,attr"`
   
   Content_Protection []Content_Protection `xml:"ContentProtection"`
   Segment_Template *Segment_Template `xml:"SegmentTemplate"`
}

type Representation struct {
   Bandwidth int `xml:"bandwidth,attr"`
   Codecs string `xml:"codecs,attr"`
   Height int `xml:"height,attr"`
   ID string `xml:"id,attr"`
   Width int `xml:"width,attr"`
   Base_URL string `xml:"BaseURL"`
   Segment_Base *Segment_Base `xml:"SegmentBase"`
   
   Content_Protection []Content_Protection `xml:"ContentProtection"`
   Segment_Template *Segment_Template `xml:"SegmentTemplate"`
}
