package command

import (
  "fmt"
)

func CommandMkdir(options map[string]string, rest []string, cwd string) {
  fmt.Println(options)
}
