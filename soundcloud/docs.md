# Overview

package `soundcloud`

## Index

- [Types](#types)
  - [type ClientMedia](#type-clientmedia)
  - [type ClientTrack](#type-clienttrack)
    - [func (c \*ClientTrack) Artwork() string](#func-clienttrack-artwork)
    - [func (c \*ClientTrack) New(id int64) error](#func-clienttrack-new)
    - [func (c \*ClientTrack) Progressive() (\*Transcoding, bool)](#func-clienttrack-progressive)
    - [func (c \*ClientTrack) Resolve(address string) error](#func-clienttrack-resolve)
  - [type Transcoding](#type-transcoding)
    - [func (t \*Transcoding) Media() (\*ClientMedia, error)](#func-transcoding-media)
- [Source files](#source-files)

## Types

### type [ClientMedia](./soundcloud.go#L72)

```go
type ClientMedia struct {
  Url string // cf-media.sndcdn.com/QaV7QR1lxpc6.128.mp3
}
```

### type [ClientTrack](./soundcloud.go#L101)

```go
type ClientTrack struct {
  ArtworkUrl  string    `json:"artwork_url"`
  DisplayDate time.Time `json:"display_date"`
  Id          int64
  Media       struct {
    Transcodings []Transcoding
  }
  Title string
  User  struct {
    AvatarUrl string `json:"avatar_url"`
    Username  string
  }
}
```

### func (\*ClientTrack) [Artwork](./soundcloud.go#L77)

```go
func (c *ClientTrack) Artwork() string
```

i1.sndcdn.com/artworks-000308141235-7ep8lo-large.jpg

### func (\*ClientTrack) [New](./soundcloud.go#L46)

```go
func (c *ClientTrack) New(id int64) error
```

### func (\*ClientTrack) [Progressive](./soundcloud.go#L16)

```go
func (c *ClientTrack) Progressive() (*Transcoding, bool)
```

Also available is "hls", but all transcodings are quality "sq".
Same for "api-mobile.soundcloud.com".

### func (\*ClientTrack) [Resolve](./soundcloud.go#L84)

```go
func (c *ClientTrack) Resolve(address string) error
```

### type [Transcoding](./soundcloud.go#L115)

```go
type Transcoding struct {
  Format struct {
    Protocol string
  }
  Url string
}
```

### func (\*Transcoding) [Media](./soundcloud.go#L25)

```go
func (t *Transcoding) Media() (*ClientMedia, error)
```

## Source files

[soundcloud.go](./soundcloud.go)
