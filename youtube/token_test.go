package youtube

import (
   "os"
   "testing"
   "time"
)

func TestAndroidCheck(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   text, err := os.ReadFile(home + "/youtube.json")
   if err != nil {
      t.Fatal(err)
   }
   var token AuthToken
   err = token.Unmarshal(text)
   if err != nil {
      t.Fatal(err)
   }
   for _, check_id := range check_ids {
      var tube InnerTube
      tube.VideoId = check_id
      tube.Context.Client.ClientName = android
      play, err := tube.Player(&token)
      if err != nil {
         t.Fatal(err)
      }
      if play.PlayabilityStatus.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
}

var check_ids = []string{
   "Cr381pDsSsA", // racy check
   "HsUATh_Nc2U", // racy check
   "SZJvDhaSDnc", // racy check
   "Tq92D6wQ1mg", // racy check
   "dqRZDebPIGs", // racy check
   "nGC3D_FkCmg", // content check
}
