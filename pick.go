package main

import (
  "fmt"
  "github.com/manifoldco/promptui"
  "os"
  "os/user"
  "path"
)

func main() {
  cwd, err := os.Getwd()
  if err != nil {
    panic(err)
  }

  usr, _ := user.Current()
  logFile := path.Join(usr.HomeDir, ".myhist.log")
  hist := history(logFile)
  items := predict(hist, cwd)
  prompt := promptui.Select {
    Label: "Pick next command",
    Items: items,
    Size: 10,
  }

  _, result, err := prompt.Run()
  if err == nil {
    fmt.Fprintf(os.Stderr, "%s\n", result)
  }
}
