package youtube

import (
   "154.pages.dev/text"
   "errors"
   "fmt"
   "mime"
   "net/url"
   "path"
   "strings"
)

type VideoId string

func (v *VideoId) Set(text string) error {
   address, err := url.Parse(text)
   if err != nil {
      return err
   }
   id := address.Query().Get("v")
   if id == "" {
      id = path.Base(address.Path)
   }
   *v = VideoId(id)
   return nil
}

func (v VideoId) String() string {
   return string(v)
}

type YtImg struct {
   Height int
   Name string
   VideoId string
   Width int
}

var YtImgs = []YtImg{
   {Width:120, Height:90, Name:"default.jpg"},
   {Width:120, Height:90, Name:"1.jpg"},
   {Width:120, Height:90, Name:"2.jpg"},
   {Width:120, Height:90, Name:"3.jpg"},
   {Width:120, Height:90, Name:"default.webp"},
   {Width:120, Height:90, Name:"1.webp"},
   {Width:120, Height:90, Name:"2.webp"},
   {Width:120, Height:90, Name:"3.webp"},
   {Width:320, Height:180, Name:"mq1.jpg"},
   {Width:320, Height:180, Name:"mq2.jpg"},
   {Width:320, Height:180, Name:"mq3.jpg"},
   {Width:320, Height:180, Name:"mqdefault.jpg"},
   {Width:320, Height:180, Name:"mq1.webp"},
   {Width:320, Height:180, Name:"mq2.webp"},
   {Width:320, Height:180, Name:"mq3.webp"},
   {Width:320, Height:180, Name:"mqdefault.webp"},
   {Width:480, Height:360, Name:"0.jpg"},
   {Width:480, Height:360, Name:"hqdefault.jpg"},
   {Width:480, Height:360, Name:"hq1.jpg"},
   {Width:480, Height:360, Name:"hq2.jpg"},
   {Width:480, Height:360, Name:"hq3.jpg"},
   {Width:480, Height:360, Name:"0.webp"},
   {Width:480, Height:360, Name:"hqdefault.webp"},
   {Width:480, Height:360, Name:"hq1.webp"},
   {Width:480, Height:360, Name:"hq2.webp"},
   {Width:480, Height:360, Name:"hq3.webp"},
   {Width:640, Height:480, Name:"sddefault.jpg"},
   {Width:640, Height:480, Name:"sd1.jpg"},
   {Width:640, Height:480, Name:"sd2.jpg"},
   {Width:640, Height:480, Name:"sd3.jpg"},
   {Width:640, Height:480, Name:"sddefault.webp"},
   {Width:640, Height:480, Name:"sd1.webp"},
   {Width:640, Height:480, Name:"sd2.webp"},
   {Width:640, Height:480, Name:"sd3.webp"},
   {Width:1280, Height:720, Name:"hq720.jpg"},
   {Width:1280, Height:720, Name:"maxresdefault.jpg"},
   {Width:1280, Height:720, Name:"maxres1.jpg"},
   {Width:1280, Height:720, Name:"maxres2.jpg"},
   {Width:1280, Height:720, Name:"maxres3.jpg"},
   {Width:1280, Height:720, Name:"hq720.webp"},
   {Width:1280, Height:720, Name:"maxresdefault.webp"},
   {Width:1280, Height:720, Name:"maxres1.webp"},
   {Width:1280, Height:720, Name:"maxres2.webp"},
   {Width:1280, Height:720, Name:"maxres3.webp"},
}

func (y *YtImg) String() string {
   var b strings.Builder
   b.WriteString("http://i.ytimg.com/vi")
   if strings.HasSuffix(y.Name, ".webp") {
      b.WriteString("_webp")
   }
   b.WriteByte('/')
   b.WriteString(y.VideoId)
   b.WriteByte('/')
   b.WriteString(y.Name)
   return b.String()
}

const (
   android_version = "19.33.35"
   web_version = "2.20231219.04.00"
)

const (
   android = "ANDROID"
   android_embedded_player = "ANDROID_EMBEDDED_PLAYER"
   web = "WEB"
)

var ClientName = []string{
   android,
   android_embedded_player,
   web,
}

// need `osVersion` this to get the correct:
// This video requires payment to watch
// instead of the invalid:
// This video can only be played on newer versions of Android or other
// supported devices.
type InnerTube struct {
   ContentCheckOk bool `json:"contentCheckOk,omitempty"`
   Context struct {
      Client struct {
         AndroidSdkVersion int `json:"androidSdkVersion"`
         ClientName string `json:"clientName"`
         ClientVersion string `json:"clientVersion"`
         OsVersion string `json:"osVersion"`
      } `json:"client"`
   } `json:"context"`
   RacyCheckOk bool `json:"racyCheckOk,omitempty"`
   VideoId string `json:"videoId"`
}

func (a *AdaptiveFormat) Ext() (string, error) {
   media, _, err := mime.ParseMediaType(a.MimeType)
   if err != nil {
      return "", err
   }
   switch media {
   case "audio/mp4":
      return ".m4a", nil
   case "audio/webm":
      return ".weba", nil
   case "video/mp4":
      return ".m4v", nil
   case "video/webm":
      return ".webm", nil
   }
   return "", errors.New(a.MimeType)
}

func (a AdaptiveFormat) CompareBitrate(b AdaptiveFormat) int {
   return a.Bitrate - b.Bitrate
}

type AdaptiveFormat struct {
   AudioQuality string
   Bitrate int
   ContentLength int64 `json:",string"`
   Itag int
   MimeType string
   QualityLabel string
   Url string
}

// we need the length for progress meter, so cannot use a channel
func (a *AdaptiveFormat) Ranges() []string {
   const bytes = 10_000_000
   var (
      byte_ranges []string
      pos int64
   )
   for pos < a.ContentLength {
      byte_range := fmt.Sprintf("&range=%v-%v", pos, pos+bytes-1)
      byte_ranges = append(byte_ranges, byte_range)
      pos += bytes
   }
   return byte_ranges
}

func (a *AdaptiveFormat) String() string {
   var b []byte
   b = fmt.Append(b, "itag = ", a.Itag)
   if a.QualityLabel != "" {
      b = fmt.Append(b, "\nlabel = ", a.QualityLabel)
   }
   b = fmt.Append(b, "\nrate = ", text.Rate(a.Bitrate))
   b = fmt.Append(b, "\nsize = ", text.Size(a.ContentLength))
   b = fmt.Append(b, "\ntype = ", a.MimeType)
   if a.AudioQuality != "" {
      b = fmt.Append(b, "\naudio = ", a.AudioQuality)
   }
   return string(b)
}
