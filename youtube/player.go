package youtube

import (
   "bytes"
   "encoding/json"
   "net/http"
   "time"
)

func (p *Player) Post(r Request, auth *OAuth) error {
   r.Context.Client.AndroidSdkVersion = 32
   r.Context.Client.OsVersion = "12"
   body, err := json.Marshal(r)
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/player",
      bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.Header.Set("User-Agent", user_agent + r.Context.Client.ClientVersion)
   if auth != nil {
      req.Header.Set("Authorization", "Bearer " + auth.V.Access_Token)
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(p)
}

func (p Player) Format(itag int) (*AdaptiveFormat, bool) {
   for _, format := range p.StreamingData.AdaptiveFormats {
      if format.Itag == itag {
         return &format, true
      }
   }
   return nil, false
}

type Player struct {
   Microformat struct {
      PlayerMicroformatRenderer struct {
         PublishDate string
      }
   }
   PlayabilityStatus struct {
      Status string
      Reason string
   }
   StreamingData struct {
      AdaptiveFormats []AdaptiveFormat
   }
   VideoDetails struct {
      Author string
      LengthSeconds int64 `json:",string"`
      ShortDescription string
      Title string
      VideoId string
      ViewCount int64 `json:",string"`
   }
}

const user_agent = "com.google.android.youtube/"

func (p Player) Author() string {
   return p.VideoDetails.Author
}

func (p Player) Time() (time.Time, error) {
   return time.Parse(
      time.RFC3339, p.Microformat.PlayerMicroformatRenderer.PublishDate,
   )
}

func (p Player) Title() string {
   return p.VideoDetails.Title
}
