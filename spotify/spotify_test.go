package spotify

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func TestOk(t *testing.T) {
   time.Sleep(9 * time.Second)
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
   login, err := response.ok(username, password)
   if err != nil {
      t.Fatal(err)
   }
   token, ok := login.access_token()
   fmt.Printf("%q %v\n", token, ok)
}
