package command

import (
  "fmt"
)

/**
 * remove files or directories
 * @param {map[string]string} options
 *    command options
 * @param {[]string} rest
 *    rest arguments
 * @param {string} cwd
 *    current working directory
 * @return {error}
 *
 */
func CommandRemove(options map[string]string, rest []string, cwd string) {
  fmt.Println(options)
}
