# spotify

~~~
"canonical_uri": "spotify:track:1oaaSrDJimABpOdCEbw2DJ"
"file_id": "98b53c239db400b0a016145700de431f68d28f54",
"gid": "2da9a11032664413b24de181c534f157",
~~~

https://play.google.com/store/apps/details?id=com.spotify.music

~~~
> play -a com.spotify.music
files: APK APK APK APK
name: Spotify: Music and Podcasts
offered by: Spotify AB
price: 0 USD
requires: 5.0 and up
size: 72.02 megabyte
updated on: Feb 26, 2024
version code: 111414784
version name: 8.9.18.512
~~~

Create Android 6 device. Install user certificate.

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://open.spotify.com/track/1oaaSrDJimABpOdCEbw2DJ
~~~

- https://github.com/Ahmeth4n/javatify/issues/3
- https://github.com/glomatico/spotify-aac-downloader/issues/21

## formats

both blocks:

~~~
protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\x98\xb5<#\x9d\xb4\x00\xb0\xa0\x16\x14W\x00\xdeC\x1fhҏT")},
protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(2)},
OGG_VORBIS_320 "\x98\xb5<#\x9d\xb4\x00\xb0\xa0\x16\x14W\x00\xdeC\x1fhҏT"

protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\xf6\x82ҩ]\x0e\x14\xee\xefO@\xb6\x0f\xdd\xdeV\xbcg!\xc7")},
protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(0)},
OGG_VORBIS_96 "\xf6\x82ҩ]\x0e\x14\xee\xefO@\xb6\x0f\xdd\xdeV\xbcg!\xc7"

protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("j_\x12\xfaQ\xf2\xc1ℯpz\x99\xf3ʆ\x96\xf7\xf6/")},
protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(1)},
OGG_VORBIS_160 "j_\x12\xfaQ\xf2\xc1ℯpz\x99\xf3ʆ\x96\xf7\xf6/"

protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\x06J\xff\v\xf7'8[\xcby\xe0\xa3֕֔\xd8\xc0,\xfe")},
protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(8)},
AAC_24 "\x06J\xff\v\xf7'8[\xcby\xe0\xa3֕֔\xd8\xc0,\xfe"
~~~

second block path, 6 steps:

~~~go
protobuf.Field{Number: 2, Type: -2, Value: protobuf.Message{
   protobuf.Field{Number: 3, Type: -2, Value: protobuf.Message{
      protobuf.Field{Number: 3, Type: -2, Value: protobuf.Message{
         protobuf.Field{Number: 2, Type: -2, Value: protobuf.Message{
            protobuf.Field{Number: 12, Type: -2, Value: protobuf.Message{
               protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\x98\xb5<#\x9d\xb4\x00\xb0\xa0\x16\x14W\x00\xdeC\x1fhҏT")},
~~~

second block only:

~~~go
protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("9$\x82\xfe\x9b\xedsr\xd1e}~\"\xf3+y)\x02\xf3\xbd")},
protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(10)},
MP4_128 "9$\x82\xfe\x9b\xedsr\xd1e}~\"\xf3+y)\x02\xf3\xbd"

protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("l\x9ex0\xa3\xf4\xbcw\xd08K\xc3\xf5\xd2}\xf5 ^Q0")},
protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(11)},
MP4_256 "l\x9ex0\xa3\xf4\xbcw\xd08K\xc3\xf5\xd2}\xf5 ^Q0"

protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\x93\xd1v\\\xfd\xc5\xe8JA`\x00Qs\x1az\xffʢ\xe9\x12")},
protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(12)},
MP4_128_DUAL "\x93\xd1v\\\xfd\xc5\xe8JA`\x00Qs\x1az\xffʢ\xe9\x12"

protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\x17=\xf6\x8bP\x97\xd4\a\xde\xff\x84\x1fAWl\x81\xc3\x0e\x13B")},
protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(13)},
MP4_256_DUAL "\x17=\xf6\x8bP\x97\xd4\a\xde\xff\x84\x1fAWl\x81\xc3\x0e\x13B"
~~~

first block path, 7 steps:

~~~go
protobuf.Field{Number: 2, Type: -2, Value: protobuf.Message{
   protobuf.Field{Number: 3, Type: -2, Value: protobuf.Message{
      protobuf.Field{Number: 3, Type: -2, Value: protobuf.Message{
         protobuf.Field{Number: 2, Type: -2, Value: protobuf.Message{
            protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
               protobuf.Field{Number: 1, Type: -2, Value: protobuf.Message{
                  protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\x98\xb5<#\x9d\xb4\x00\xb0\xa0\x16\x14W\x00\xdeC\x1fhҏT")},
~~~

first block only:

~~~go
protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\xfe\xb3\xd0e\x0e\xc8|\xcf\xdf\x18\xb6\x89i\xb8_\xa4\xbb\x10\xfb\xb4")},
protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(16)},
FLAC_FLAC = 16;

protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("{T\x91\xd2rO\xf2HG\x16\x19.A\xe8\x06v\x824\xf68")},
protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(18)},

protobuf.Field{Number: 1, Type: 2, Value: protobuf.Bytes("\xbd\x03z\xb3\x18a\x053\xfbk\xc5 \xe2:!\x83\x18\xf6C\xa8")},
protobuf.Field{Number: 2, Type: 0, Value: protobuf.Varint(19)},
~~~
