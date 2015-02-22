package main

import (
  "github.com/chrisn-au/go-limitless"
)

func main() {
  c := limitless.LimitlessController{}
  c.Host = "192.168.0.25"
  group := limitless.LimitlessGroup{}
  group.Id = 2
  group.Controller = &c
  c.Groups = []limitless.LimitlessGroup{group}

  group.On()
}
