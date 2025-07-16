package moonshot

import (
   "bytes"
   "encoding/json"
   "net/http"
)

func completions(bearer string, messages []message) (*http.Response, error) {
   var value struct {
      Model string `json:"model"`
      Messages []message `json:"messages"`
   }
   value.Model = "moonshot-v1-8k"
   value.Messages = messages
   data, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://api.moonshot.ai/v1/chat/completions",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("authorization", "Bearer " + bearer)
   return http.DefaultClient.Do(req)
}

type content struct {
   Content string `json:"content"`
   FileName string `json:"file_name"`
}

type message struct {
   Role string `json:"role"`
   Content string `json:"content"`
}

func caching(bearer string, contents []content) (*http.Response, error) {
   var value struct {
      Model string `json:"model"`
      Messages []message `json:"messages"`
   }
   value.Model = "moonshot-v1"
   value.Messages = make([]message, len(contents))
   for i, contentVar := range contents {
      data, err := json.Marshal(contentVar)
      if err != nil {
         return nil, err
      }
      value.Messages[i] = message{
         Role: "user",
         Content: string(data),
      }
   }
   data, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://api.moonshot.ai/v1/caching", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("authorization", "Bearer " + bearer)
   return http.DefaultClient.Do(req)
}
