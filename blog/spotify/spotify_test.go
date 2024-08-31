package spotify

import (
   "fmt"
   "os"
   "testing"
)

func TestOkRead(t *testing.T) {
   var (
      login LoginOk
      err error
   )
   login.Data, err = os.ReadFile("spotify.txt")
   if err != nil {
      t.Fatal(err)
   }
   if err := login.Consume(); err != nil {
      t.Fatal(err)
   }
   fmt.Println(login.AccessToken())
}

func TestOkWrite(t *testing.T) {
   username := os.Getenv("spotify_username")
   if username == "" {
      t.Fatal("spotify_username")
   }
   password := os.Getenv("spotify_password")
   var response login_response
   err := response.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   ok, err := response.ok(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("spotify.txt", ok.Data, 0666)
}
