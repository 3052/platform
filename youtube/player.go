package youtube

import (
   "bytes"
   "encoding/json"
   "net/http"
   "time"
)

func (i *InnerTube) Player(token *AuthToken) (*Player, error) {
   i.Context.Client.AndroidSdkVersion = 32
   i.Context.Client.OsVersion = "12"
   switch i.Context.Client.ClientName {
   case android:
      i.ContentCheckOk = true
      i.Context.Client.ClientVersion = android_version
      i.RacyCheckOk = true
   case android_embedded_player:
      i.Context.Client.ClientVersion = android_version
   case web:
      i.Context.Client.ClientVersion = web_version
   }
   body, err := json.MarshalIndent(i, "", " ")
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.youtube.com/youtubei/v1/player",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("user-agent", user_agent + i.Context.Client.ClientVersion)
   if token != nil {
      req.Header.Set("authorization", "Bearer " + token.AccessToken)
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   play := &Player{}
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}
type Player struct {
   Microformat struct {
      PlayerMicroformatRenderer struct {
         PublishDate Date
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

type Date struct {
   Time time.Time
}

func (d *Date) UnmarshalText(text []byte) error {
   var err error
   d.Time, err = time.Parse(time.RFC3339, string(text))
   if err != nil {
      return err
   }
   return nil
}
