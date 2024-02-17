package youtube

import (
   "bytes"
   "encoding/json"
   "net/http"
   "strings"
)

func (w WatchNext) Year() (string, bool) {
   if v, ok := w.metadata_row_container(); ok {
      return v.get("Release date")
   }
   return "", false
}

type WatchNext []struct {
   VideoPrimaryInfoRenderer *struct {
      Title Run // The Family Secret
   }
   VideoSecondaryInfoRenderer *struct {
      MetadataRowContainer MetadataRowContainer
      Owner struct {
         VideoOwnerRenderer struct {
            Title Run
         }
      }
   }
}

func (w WatchNext) Episode() (string, bool) {
   if v, ok := w.metadata_row_container(); ok {
      return v.get("Episode")
   }
   return "", false
}

func (w WatchNext) Owner() (string, bool) {
   for _, v := range w {
      if v := v.VideoSecondaryInfoRenderer; v != nil {
         return v.Owner.VideoOwnerRenderer.Title.String(), true
      }
   }
   return "", false
}

func (w WatchNext) Season() (string, bool) {
   if v, ok := w.metadata_row_container(); ok {
      return v.get("Season")
   }
   return "", false
}

func (w WatchNext) Show() (string, bool) {
   if v, ok := w.metadata_row_container(); ok {
      return v.get("Show")
   }
   return "", false
}

func (w WatchNext) Title() (string, bool) {
   for _, v := range w {
      if v := v.VideoPrimaryInfoRenderer; v != nil {
         return v.Title.String(), true
      }
   }
   return "", false
}

func (w WatchNext) metadata_row_container() (*MetadataRowContainer, bool) {
   for _, v := range w {
      if v := v.VideoSecondaryInfoRenderer; v != nil {
         return &v.MetadataRowContainer, true
      }
   }
   return nil, false
}

type MetadataRowContainer struct {
   MetadataRowContainerRenderer struct {
      Rows []struct {
         MetadataRowRenderer struct {
            Title Run // Show
            Contents Runs // In The Heat Of The Night
         }
      }
   }
}

func (m MetadataRowContainer) get(s string) (string, bool) {
   for _, v := range m.MetadataRowContainerRenderer.Rows {
      if v := v.MetadataRowRenderer; v.Title.String() == s {
         return v.Contents.String(), true
      }
   }
   return "", false
}

type Run struct {
   Runs []struct {
      Text string
   }
   SimpleText string
}

func (r Run) String() string {
   if r.SimpleText != "" {
      return r.SimpleText
   }
   var b strings.Builder
   for _, each := range r.Runs {
      b.WriteString(each.Text)
   }
   return b.String()
}

type Runs []Run

func (r Runs) String() string {
   var b strings.Builder
   for _, each := range r {
      if b.Len() >= 1 {
         b.WriteString(", ")
      }
      b.WriteString(each.String())
   }
   return b.String()
}

// /youtubei/v1/player is missing the name of the series. use with WEB client.
func (w *WatchNext) Post(r Request) error {
   body, err := json.Marshal(r)
   if err != nil {
      return err
   }
   res, err := http.Post(
      "https://www.youtube.com/youtubei/v1/next?prettyPrint=false" ,
      "application/json", bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   var s struct {
      Contents struct {
         TwoColumnWatchNextResults struct {
            Results struct {
               Results struct {
                  Contents WatchNext
               }
            }
         }
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return err
   }
   *w = s.Contents.TwoColumnWatchNextResults.Results.Results.Contents
   return nil
}
