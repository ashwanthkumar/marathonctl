package main

import (
  "os"
  "fmt"

  "github.com/ashwanthkumar/marathonctl/cmd"
)

func main() {
  setupSignalHandlers()

  if err := cmd.MarathonCtl.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
}
