package command

import (
  "fmt"
)

/**
 * List a directory
 * @param {map[string]string} options
 *    command options
 * @param {[]string} rest
 *    rest arguments
 * @param {string} cwd
 *    current working directory
 * @return {error}
 *
 */
func CommandList(options map[string]string, rest []string, cwd string) {
  fmt.Println(options, rest)
}
