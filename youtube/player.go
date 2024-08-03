package youtube

import (
   "bytes"
   "encoding/json"
   "net/http"
   "time"
)

func (p *Player) New(r Request, token *AuthToken) error {
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
   req.Header.Set("user-agent", user_agent + r.Context.Client.ClientVersion)
   if token != nil {
      req.Header.Set("authorization", "Bearer " + token.AccessToken)
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(p)
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
