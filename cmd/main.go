package main

import (
  "log"

  "gitlab.com/croc3/umbrellacorptask/src/routes"
)

func main() {
  r, err := routes.NewRoute()
  if err != nil {
    log.Fatal("Failed to create a new Route")
  }

  r.Process()
}
