package youtube

import (
   "os"
   "testing"
   "time"
)

var check_ids = []string{
   "Cr381pDsSsA", // racy check
   "HsUATh_Nc2U", // racy check
   "SZJvDhaSDnc", // racy check
   "Tq92D6wQ1mg", // racy check
   "dqRZDebPIGs", // racy check
   "nGC3D_FkCmg", // content check
}

func TestAndroidCheck(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   text, err := os.ReadFile(home + "/youtube.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, check_id := range check_ids {
      var play Player
      req := Request{VideoId: check_id}
      req.AndroidCheck()
      token := Token{Raw: text}
      token.Unmarshal()
      err := play.Post(req, &token)
      if err != nil {
         t.Fatal(err)
      }
      if play.PlayabilityStatus.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
}
