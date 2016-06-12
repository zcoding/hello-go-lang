package command

import (
  "fmt"
)

var version string = "0.0.1"

func CommandVersion(options map[string]string, rest []string, cwd string) {
  fmt.Println("vfs version: ", version)
}
