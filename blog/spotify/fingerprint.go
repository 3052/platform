package main

import (
   "encoding/base64"
   "fmt"
   "github.com/refraction-networking/utls"
)

const raw = "FgMDAKkBAAClAwNMBj+jgVR0HqCQa/f0QS2sKDGZJELTUVukjERQ+L1qYQAAFsArwCzAL8AwwBPAFACcAJ0ALwA1AP8BAABmAAAAHAAaAAAXc3BjbGllbnQud2cuc3BvdGlmeS5jb20AFwAAACMAAAANABYAFAYBBgMFAQUDBAEEAwMBAwMCAQIDABAADgAMAmgyCGh0dHAvMS4xAAsAAgEAAAoACAAGABcAGAAZ"

func main() {
   data, err := base64.StdEncoding.DecodeString(raw)
   if err != nil {
      panic(err)
   }
   var finger tls.Fingerprinter
   spec, err := finger.FingerprintClientHello(data)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%#v\n", spec)
}
