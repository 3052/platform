package main

import (
   "41.neocities.org/platform/nord"
   "flag"
   "fmt"
   "log"
   "os"
)

func do_country(code string) error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   data, err := os.ReadFile(home + name)
   if err != nil {
      return err
   }
   var loads nord.ServerLoads
   err = loads.Unmarshal(data)
   if err != nil {
      return err
   }
   fmt.Println(
      nord.Proxy(loads.Country(code)),
   )
   data, err = loads.Marshal()
   if err != nil {
      return err
   }
   return write_file(home + name, data)
}

func main() {
   write := flag.Bool("w", false, "write")
   country_code := flag.String("c", "", "country code")
   switch {
   case *country_code != "":
      err := do_country(*country_code)
      if err != nil {
         panic(err)
      }
   case *write:
      err := do_write()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

const name = "/platform/nord/ServerLoads"

func do_write() error {
   servers, err := nord.GetServers(0)
   if err != nil {
      return err
   }
   data, err := nord.GetServerLoads(servers).Marshal()
   if err != nil {
      return err
   }
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   return write_file(home+name, data)
}

func write_file(name string, data []byte) error {
   log.Println("WriteFile", name)
   return os.WriteFile(name, data, os.ModePerm)
}
