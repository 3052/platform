package internal

import (
   "154.pages.dev/encoding"
   "154.pages.dev/encoding/hls"
   "154.pages.dev/log"
   "io"
   "net/http"
   "net/url"
   "os"
   "text/template"
)

// wikipedia.org/wiki/HTTP_Live_Streaming
type HttpStream struct {
   Base *url.URL
   Name encoding.Namer
}

func (h HttpStream) HLS(master hls.MasterPlaylist, index int) error {
   variant, ok := master.Index(index)
   if !ok {
      line, err := new(template.Template).Parse(hls.ModeLine)
      if err != nil {
         return err
      }
      return line.Execute(os.Stdout, master)
   }
   var segment hls.MediaSegment
   req, err := http.NewRequest("", variant.URI, nil)
   if err != nil {
      return err
   }
   req.URL = h.Base.ResolveReference(req.URL)
   err = func() error {
      res, err := http.DefaultClient.Do(req)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      text, err := io.ReadAll(res.Body)
      if err != nil {
         return err
      }
      segment.New(string(text))
      return nil
   }()
   if err != nil {
      return err
   }
   var mode hls.BlockMode
   err = func() error {
      res, err := http.Get(segment.Key.URI)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      key, err := io.ReadAll(res.Body)
      if err != nil {
         return err
      }
      return mode.New(key)
   }()
   if err != nil {
      return err
   }
   name, err := encoding.Name(h.Name)
   if err != nil {
      return err
   }
   file, err := os.Create(name + variant.Ext())
   if err != nil {
      return err
   }
   defer file.Close()
   var meter log.ProgressMeter
   meter.Set(len(segment.URI))
   log.SetTransport(nil)
   defer log.Transport{}.Set()
   for _, uri := range segment.URI {
      // with HLS, the segment URL is relative to the master URL, and the
      // fragment URL is relative to the segment URL.
      req.URL, err = req.URL.Parse(uri)
      if err != nil {
         return err
      }
      err := func() error {
         res, err := http.DefaultClient.Do(req)
         if err != nil {
            return err
         }
         defer res.Body.Close()
         data, err := io.ReadAll(meter.Reader(res))
         if err != nil {
            return err
         }
         data = mode.Decrypt(data)
         if _, err := file.Write(data); err != nil {
            return err
         }
         return nil
      }()
      if err != nil {
         return err
      }
   }
   return nil
}

func (h *HttpStream) HlsMaster(uri string) (hls.MasterPlaylist, error) {
   res, err := http.Get(uri)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   h.Base = res.Request.URL
   text, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   var master hls.MasterPlaylist
   master.New(string(text))
   return master, nil
}
