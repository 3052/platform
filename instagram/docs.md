# Overview

package `instagram`

## Index

- [Types](#types)
  - [type Address](#type-address)
    - [func (a \*Address) Media() (\*ShortcodeMedia, error)](#func-address-media)
    - [func (a \*Address) Set(text string) error](#func-address-set)
    - [func (a \*Address) String() string](#func-address-string)
  - [type ShortcodeMedia](#type-shortcodemedia)
    - [func (s \*ShortcodeMedia) DisplayUrls() []Url](#func-shortcodemedia-displayurls)
  - [type Url](#type-url)
    - [func (b \*Url) UnmarshalText(text []byte) error](#func-url-unmarshaltext)
- [Source files](#source-files)

## Types

### type [Address](./instagram.go#L41)

```go
type Address struct {
  Shortcode string
}
```

### func (\*Address) [Media](./instagram.go#L45)

```go
func (a *Address) Media() (*ShortcodeMedia, error)
```

### func (\*Address) [Set](./instagram.go#L32)

```go
func (a *Address) Set(text string) error
```

### func (\*Address) [String](./instagram.go#L37)

```go
func (a *Address) String() string
```

### type [ShortcodeMedia](./instagram.go#L78)

```go
type ShortcodeMedia struct {
  DisplayUrl            Url `json:"display_url"`
  EdgeSidecarToChildren *struct {
    Edges []struct {
      Node struct {
        DisplayUrl Url `json:"display_url"`
      }
    }
  } `json:"edge_sidecar_to_children"`
}
```

### func (\*ShortcodeMedia) [DisplayUrls](./instagram.go#L11)

```go
func (s *ShortcodeMedia) DisplayUrls() []Url
```

### type [Url](./instagram.go#L22)

```go
type Url struct {
  Url url.URL
}
```

### func (\*Url) [UnmarshalText](./instagram.go#L26)

```go
func (b *Url) UnmarshalText(text []byte) error
```

## Source files

[instagram.go](./instagram.go)
