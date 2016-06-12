package command

import (
  "fmt"
)

/**
 *
 * Remove file
 * -f force
 * -r recursive
 * -v verbose
 * --dry-run dry run and show results, but do nothing
 *
 */
func CommandRemove(options map[string]string, rest []string, cwd string) {
  fmt.Println(options)
}
