package main

import (
  "bufio"
  "os"
  "regexp"
  "strings"
)

type HistEntry struct {
  Cwd string
  Cmd string
}

func history(histFile string) []HistEntry {
  f, err := os.Open(histFile)
  if err != nil {
    panic(err)
  }
  defer f.Close()
  hist := []HistEntry{}
  scanner := bufio.NewScanner(f)
  cmdSane := regexp.MustCompile(`^\S`)
  for scanner.Scan() {
    // epoch cwd cmd
    flds := strings.SplitN(scanner.Text(), " ", 3)
    if len(flds) != 3 || !cmdSane.MatchString(flds[2]) || flds[2] == "g" {
      continue
    }
    hist = append(hist, HistEntry{ Cwd: flds[1], Cmd: flds[2] })
  }

  if err := scanner.Err(); err != nil {
    panic(err)
  }

  return hist
}
