# Overview

package `bandcamp`

## Index

- [Constants](#constants)
- [Variables](#variables)
- [Types](#types)
  - [type AlbumTrack](#type-albumtrack)
  - [type BandDetails](#type-banddetails)
    - [func (b \*BandDetails) New(id int64) error](#func-banddetails-new)
  - [type Image](#type-image)
    - [func (i \*Image) URL(art_id int64) string](#func-image-url)
  - [type Item](#type-item)
    - [func (i \*Item) Band() (\*BandDetails, error)](#func-item-band)
    - [func (i \*Item) Tralbum() (\*Tralbum, error)](#func-item-tralbum)
  - [type ReportParams](#type-reportparams)
    - [func (r \*ReportParams) Band() (\*BandDetails, error)](#func-reportparams-band)
    - [func (r \*ReportParams) New(address string) error](#func-reportparams-new)
    - [func (r \*ReportParams) Tralbum() (\*Tralbum, error)](#func-reportparams-tralbum)
  - [type Tralbum](#type-tralbum)
    - [func (t \*Tralbum) Date() time.Time](#func-tralbum-date)
- [Source files](#source-files)

## Constants

```go
const (
  Jpeg = iota
  Png
)
```

## Variables

```go
var Images = []Image{
  {ID: 0, Width: 1500, Height: 1500, Format: Jpeg},
  {ID: 1, Width: 1500, Height: 1500, Format: Png},
  {ID: 2, Width: 350, Height: 350, Format: Jpeg},
  {ID: 3, Width: 100, Height: 100, Format: Jpeg},
  {ID: 4, Width: 300, Height: 300, Format: Jpeg},
  {ID: 5, Width: 700, Height: 700, Format: Jpeg},
  {ID: 6, Width: 100, Height: 100, Format: Jpeg},
  {ID: 7, Width: 150, Height: 150, Format: Jpeg},
  {ID: 8, Width: 124, Height: 124, Format: Jpeg},
  {ID: 9, Width: 210, Height: 210, Format: Jpeg},
  {ID: 10, Width: 1200, Height: 1200, Format: Jpeg},
  {ID: 11, Width: 172, Height: 172, Format: Jpeg},
  {ID: 12, Width: 138, Height: 138, Format: Jpeg},
  {ID: 13, Width: 380, Height: 380, Format: Jpeg},
  {ID: 14, Width: 368, Height: 368, Format: Jpeg},
  {ID: 15, Width: 135, Height: 135, Format: Jpeg},
  {ID: 16, Width: 700, Height: 700, Format: Jpeg},
  {ID: 20, Width: 1024, Height: 1024, Format: Jpeg},
  {ID: 21, Width: 120, Height: 120, Format: Jpeg},
  {ID: 22, Width: 25, Height: 25, Format: Jpeg},
  {ID: 23, Width: 300, Height: 300, Format: Jpeg},
  {ID: 24, Width: 300, Height: 300, Format: Jpeg},
  {ID: 25, Width: 700, Height: 700, Format: Jpeg},
  {ID: 26, Width: 800, Height: 600, Format: Jpeg, Crop: true},
  {ID: 27, Width: 715, Height: 402, Format: Jpeg, Crop: true},
  {ID: 28, Width: 768, Height: 432, Format: Jpeg, Crop: true},
  {ID: 29, Width: 100, Height: 75, Format: Jpeg, Crop: true},
  {ID: 31, Width: 1024, Height: 1024, Format: Png},
  {ID: 32, Width: 380, Height: 285, Format: Jpeg, Crop: true},
  {ID: 33, Width: 368, Height: 276, Format: Jpeg, Crop: true},
  {ID: 36, Width: 400, Height: 300, Format: Jpeg, Crop: true},
  {ID: 37, Width: 168, Height: 126, Format: Jpeg, Crop: true},
  {ID: 38, Width: 144, Height: 108, Format: Jpeg, Crop: true},
  {ID: 41, Width: 210, Height: 210, Format: Jpeg},
  {ID: 42, Width: 50, Height: 50, Format: Jpeg},
  {ID: 43, Width: 100, Height: 100, Format: Jpeg},
  {ID: 44, Width: 200, Height: 200, Format: Jpeg},
  {ID: 50, Width: 140, Height: 140, Format: Jpeg},
  {ID: 65, Width: 700, Height: 700, Format: Jpeg},
  {ID: 66, Width: 1200, Height: 1200, Format: Jpeg},
  {ID: 67, Width: 350, Height: 350, Format: Jpeg},
  {ID: 68, Width: 210, Height: 210, Format: Jpeg},
  {ID: 69, Width: 700, Height: 700, Format: Jpeg},
}
```

## Types

### type [AlbumTrack](./bandcamp.go#L95)

```go
type AlbumTrack struct {
  TrackNum     int64 `json:"track_num"`
  Title        string
  BandName     string `json:"band_name"`
  StreamingUrl *struct {
    MP3_128 string `json:"mp3-128"`
  } `json:"streaming_url"`
}
```

### type [BandDetails](./bandcamp.go#L90)

```go
type BandDetails struct {
  Name        string
  Discography []Item
}
```

### func (\*BandDetails) [New](./bandcamp.go#L14)

```go
func (b *BandDetails) New(id int64) error
```

### type [Image](./bandcamp.go#L131)

```go
type Image struct {
  Crop   bool
  Format int
  Height int
  ID     int64
  Width  int
}
```

### func (\*Image) [URL](./bandcamp.go#L186)

```go
func (i *Image) URL(art_id int64) string
```

Extension is optional.

### type [Item](./bandcamp.go#L116)

```go
type Item struct {
  BandId   int64  `json:"band_id"`
  ItemId   int    `json:"item_id"`
  ItemType string `json:"item_type"`
}
```

### func (\*Item) [Band](./bandcamp.go#L122)

```go
func (i *Item) Band() (*BandDetails, error)
```

### func (\*Item) [Tralbum](./bandcamp.go#L195)

```go
func (i *Item) Tralbum() (*Tralbum, error)
```

### type [ReportParams](./bandcamp.go#L58)

```go
type ReportParams struct {
  Aid   int64  `json:"a_id"`
  Iid   int    `json:"i_id"`
  Itype string `json:"i_type"`
}
```

### func (\*ReportParams) [Band](./bandcamp.go#L49)

```go
func (r *ReportParams) Band() (*BandDetails, error)
```

### func (\*ReportParams) [New](./bandcamp.go#L28)

```go
func (r *ReportParams) New(address string) error
```

### func (\*ReportParams) [Tralbum](./bandcamp.go#L64)

```go
func (r *ReportParams) Tralbum() (*Tralbum, error)
```

### type [Tralbum](./bandcamp.go#L108)

```go
type Tralbum struct {
  ArtId         int64 `json:"art_id"`
  Title         string
  Tracks        []AlbumTrack
  ReleaseDate   int64  `json:"release_date"`
  TralbumArtist string `json:"tralbum_artist"`
}
```

### func (\*Tralbum) [Date](./bandcamp.go#L104)

```go
func (t *Tralbum) Date() time.Time
```

## Source files

[bandcamp.go](./bandcamp.go)
