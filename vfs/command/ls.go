package command

import (
  "fmt"
)

/**
 *
 * List files
 * -l long format
 * -a all files
 * -h human readable
 * -t sort by modify time
 *
 */
func CommandList(options map[string]string, rest []string, cwd string) {
  fmt.Println(options, rest)
}
