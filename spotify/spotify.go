package spotify

import (
   "154.pages.dev/protobuf"
   "bytes"
   "crypto/sha1"
   "encoding/binary"
   "errors"
   "io"
   "math/bits"
   "net/http"
)

const android_client_id = "9a8d2f0ce77a4e248bb71fefcb557637"

// github.com/librespot-org/librespot/blob/dev/core/src/spclient.rs
func solve_hash_cash(login_context, prefix []byte, length int) []byte {
   sum := sha1.Sum(login_context)
   var counter uint64
   target := binary.BigEndian.Uint64(sum[12:])
   for {
      suffix := func() []byte {
         b := binary.BigEndian.AppendUint64(nil, target)
         return binary.BigEndian.AppendUint64(b, counter)
      }()
      sum := sha1.Sum(append(prefix, suffix...))
      if bits.TrailingZeros64(binary.BigEndian.Uint64(sum[12:])) >= length {
         return suffix
      }
      counter++
      target++
   }
}

// github.com/librespot-org/librespot/blob/dev/protocol/proto/spotify/login5/v3/login5.proto
type login_ok struct {
   message protobuf.Message
}

// we are converting the option to result multiple times, so lets go ahead and
// return a result
func (o *login_ok) access_token() (string, bool) {
   m, _ := o.message.Get(1)() // LoginOk ok
   v, ok := m.GetBytes(2)()   // string access_token
   return string(v), ok
}

type login_response struct {
   message protobuf.Message
}

func (r *login_response) ok(username, password string) (*login_ok, error) {
   login_context, ok := r.login_context()
   if !ok {
      return nil, errors.New("login_response.login_context")
   }
   prefix, ok := r.prefix()
   if !ok {
      return nil, errors.New("login_response.prefix")
   }
   m := protobuf.Message{}
   m.Add(1, func(m protobuf.Message) {
      m.AddBytes(1, []byte(android_client_id))
   })
   m.AddBytes(2, login_context)
   m.Add(3, func(m protobuf.Message) {
      m.Add(1, func(m protobuf.Message) {
         m.Add(1, func(m protobuf.Message) {
            m.AddBytes(1, solve_hash_cash(login_context, prefix, 10))
         })
      })
   })
   m.Add(101, func(m protobuf.Message) {
      m.AddBytes(1, []byte(username))
      m.AddBytes(2, []byte(password))
   })
   req, err := http.NewRequest(
      "POST", "https://login5.spotify.com/v3/login",
      bytes.NewReader(m.Marshal()),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "content-type": {"application/x-protobuf"},
      "user-agent":   {"Spotify/8.9.68.456 Android/23 (Android SDK built for x86)"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   m = protobuf.Message{}
   err = m.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   return &login_ok{m}, nil
}

func (r *login_response) login_context() ([]byte, bool) {
   return r.message.GetBytes(5)()
}

// github.com/librespot-org/librespot/blob/dev/protocol/proto/spotify/login5/v3/challenges/hashcash.proto
// github.com/librespot-org/librespot/blob/dev/protocol/proto/spotify/login5/v3/login5.proto
func (r *login_response) prefix() ([]byte, bool) {
   m, _ := r.message.Get(3)() // Challenges
   m, _ = m.Get(1)()          // Challenge
   m, _ = m.Get(1)()          // challenges.HashcashChallenge
   return m.GetBytes(1)()
}

func (r *login_response) New(username, password string) error {
   m := protobuf.Message{}
   m.Add(1, func(m protobuf.Message) {
      m.AddBytes(1, []byte(android_client_id))
   })
   m.Add(101, func(m protobuf.Message) {
      m.AddBytes(1, []byte(username))
      m.AddBytes(2, []byte(password))
   })
   resp, err := http.Post(
      "https://login5.spotify.com/v3/login",
      "application/x-protobuf", bytes.NewReader(m.Marshal()),
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   r.message = protobuf.Message{}
   return r.message.Unmarshal(data)
}
