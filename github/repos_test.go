package github

import (
   "fmt"
   "os"
   "testing"
   "time"
)

var repos = []repository{
   {
      name: "media",
      description: "Download media or send API requests",
      topics: []string{
         "hulu",
         "mubi",
         "nbc",
         "paramount",
         "peacock",
         "plex",
         "stan",

         "amc",
         "roku",
      },
   },
}

const sleep = 99*time.Millisecond

func Test_Repo(t *testing.T) {
   username := os.Getenv("github_username")
   password := os.Getenv("github_password")
   for _, repo := range repos {
      fmt.Println(repo.name)
      err := repo.set_actions(username, password)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(sleep)
      if err := repo.set_description(username, password); err != nil {
         t.Fatal(err)
      }
      time.Sleep(sleep)
      if repo.topics != nil {
         err := repo.set_topics(username, password)
         if err != nil {
            t.Fatal(err)
         }
         time.Sleep(sleep)
      }
   }
}

