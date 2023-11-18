package dash

import (
   "encoding/base64"
   "encoding/hex"
   "encoding/xml"
   "errors"
   "io"
   "strconv"
   "strings"
)

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

type Segment_Base struct {
   Index_Range string `xml:"indexRange,attr"`
}

func (a Adaptation) Type() string {
   if a.MIME_Type != "" {
      return a.MIME_Type
   }
   return a.Content_Type
}

func (s Segment_Base) Start() (int64, error) {
   i := strings.Index(s.Index_Range, "-")
   return strconv.ParseInt(s.Index_Range[:i], 10, 64)
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

type Content_Protection struct {
   Default_KID string `xml:"default_KID,attr"`
   PSSH string `xml:"pssh"`
   Scheme_ID_URI string `xml:"schemeIdUri,attr"`
}

type Adaptation struct {
   Lang string `xml:"lang,attr"`
   Role *struct {
      Value string `xml:"value,attr"`
   }
   Content_Type string `xml:"contentType,attr"`
   MIME_Type string `xml:"mimeType,attr"`
   Representation []Representation
   Segment_Template *Segment_Template `xml:"SegmentTemplate"`
   Content_Protection []Content_Protection `xml:"ContentProtection"`
}

func Adaptations(r io.Reader) ([]Adaptation, error) {
   var s struct {
      Period struct {
         Adaptation_Set []Adaptation `xml:"AdaptationSet"`
      }
   }
   err := xml.NewDecoder(r).Decode(&s)
   if err != nil {
      return nil, err
   }
   return s.Period.Adaptation_Set, nil
}

type Segment_Template struct {
   Start_Number int `xml:"startNumber,attr"`
   Segment_Timeline struct {
      S []struct {
         D int `xml:"d,attr"` // duration
         R int `xml:"r,attr"` // repeat
         T int `xml:"t,attr"` // time
      }
   } `xml:"SegmentTimeline"`
   Initialization Initialization `xml:"initialization,attr"`
   Media Media `xml:"media,attr"`
}

type Initialization string

type Media string

type Representation struct {
   Bandwidth int `xml:"bandwidth,attr"`
   Codecs string `xml:"codecs,attr"`
   Height int `xml:"height,attr"`
   ID string `xml:"id,attr"`
   Width int `xml:"width,attr"`
   Base_URL string `xml:"BaseURL"`
   Segment_Base *Segment_Base `xml:"SegmentBase"`
   Segment_Template *Segment_Template `xml:"SegmentTemplate"`
   Content_Protection []Content_Protection `xml:"ContentProtection"`
}

func (i Initialization) Replace(id string) string {
   return strings.Replace(string(i), "$RepresentationID$", id, 1)
}

func (r Representation) PSSH() ([]byte, error) {
   for _, c := range r.Content_Protection {
      if c.Scheme_ID_URI == "urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed" {
         return base64.StdEncoding.DecodeString(c.PSSH)
      }
   }
   return nil, errors.New("pssh")
}

func (m Media) Replace(r Representation) []string {
   var refs []string
   for _, segment := range r.Segment_Template.Segment_Timeline.S {
      segment.T = r.Segment_Template.Start_Number
      for segment.R >= 0 {
         ref := func(s string) string {
            s = strings.Replace(s, "$Number$", strconv.Itoa(segment.T), 1)
            return strings.Replace(s, "$RepresentationID$", r.ID, 1)
         }(string(m))
         refs = append(refs, ref)
         r.Segment_Template.Start_Number++
         segment.R--
         segment.T++
      }
   }
   return refs
}

func (r Representation) Default_KID() ([]byte, error) {
   for _, c := range r.Content_Protection {
      if c.Scheme_ID_URI == "urn:mpeg:dash:mp4protection:2011" {
         c.Default_KID = strings.ReplaceAll(c.Default_KID, "-", "")
         return hex.DecodeString(c.Default_KID)
      }
   }
   return nil, errors.New("default_KID")
}
