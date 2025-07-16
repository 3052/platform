package moonshot

import (
   "os"
   "os/exec"
   "testing"
)

func TestCompletion(t *testing.T) {
   bearer, err := exec.Command("password", "-i", "moonshot.ai").Output()
   if err != nil {
      t.Fatal(err)
   }
   messages := []message{
      {
         Role: "cache",
         Content: "cache_id=cache-f3uha4moc6di11bccjbi",
      },
      {
         Role: "user",
         Content: "print second word in pricing.md",
      },
   }
   resp, err := completions(string(bearer), messages)
   if err != nil {
      t.Fatal(err)
   }
   err = resp.Write(os.Stdout)
   if err != nil {
      t.Fatal(err)
   }
}

func TestCaching(t *testing.T) {
   bearer, err := exec.Command("password", "-i", "moonshot.ai").Output()
   if err != nil {
      t.Fatal(err)
   }
   contents := []content{
      {
         Content: "the quick brown fox jumps over the lazy dog",
         FileName: "hello.txt",
      },
   }
   resp, err := caching(string(bearer), contents)
   if err != nil {
      t.Fatal(err)
   }
   err = resp.Write(os.Stdout)
   if err != nil {
      t.Fatal(err)
   }
}
