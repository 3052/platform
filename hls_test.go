package stream

import (
   "154.pages.dev/http"
   "testing"
)

const master = "https://cbcrcott-gem.akamaized.net/010605e1-250e-4a3c-a2eb-e52fbef91ead/CBC_SCORE.ism/desktop_master.m3u8?hdnea=st=1697922759~exp=1697922879~acl=/010605e1-250e-4a3c-a2eb-e52fbef91ead/cbc_score.ism/desktop*~hmac=0cd9adf908fd0363169f4a842bdbc8f5bd88ca3a62cffbeee5184e1693a0aae4"

func Test_HLS(t *testing.T) {
   http.No_Location()
   http.Verbose()
   str := Stream{Name: "Name"}
   master, err := str.HLS(master)
   if err != nil {
      t.Fatal(err)
   }
   if err := str.HLS_Streams(master.Stream, 0); err != nil {
      t.Fatal(err)
   }
}
