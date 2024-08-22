package main

import (
  "sort"
)

func predict(hist []HistEntry, cwd string) []string {
  lastCmd := ""
  followMap := map[string]map[string]int{}

  for _, h := range hist {
    if h.Cwd != cwd {
      continue
    }
    if lastCmd == "" {
      lastCmd = h.Cmd
      continue
    }
    cmdMap, ok := followMap[lastCmd]
    if !ok {
      cmdMap = map[string]int{}
      followMap[lastCmd] = cmdMap
    }
    cmdMap[h.Cmd] += 1
    lastCmd = h.Cmd
  }

  if lastCmd == "" {
    // first time in this dir
    return []string{"ls"}
  }

  items := []string{}
  follows, ok := followMap[lastCmd]

  if !ok {
    // no follow defined, just
    // return all cmds known
    for from, _ := range followMap {
      items = append(items, from)
    }

    return items
  }

  // Return best-scoring follows
  type score struct {
    to      string
    weight  int
  }

  scores := []score{}
  for to, v := range follows {
    scores = append(scores, score{to: to, weight: v})
  }

  sort.Slice(scores, func(i, j int) bool {
    return scores[i].weight > scores[j].weight
  })

  for _, score := range scores {
    items = append(items, score.to)
  }

  return items
}
