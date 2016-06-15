package command

import (
  "fmt"
  "errors"
)

/**
 * 
 * Change directory
 * @param {map[string]string} options
 *    command options
 * @param {[]string} rest
 *    rest arguments
 * @param {string} cwd
 *    current working directory
 * @return {string} destination directory
 * @return {error}
 *
 */
func CommandCd(options map[string]string, rest []string, cwd string) (string, error) {
  var destDirectory string = cwd
  destLen := len(rest)
  if destLen > 1 {
    return destDirectory, errors.New("cd: no such file or directory")
  } else if len(rest) > 0 {
    destDirectory = rest[0]
  } else {
    destDirectory = "~" // 如果没有指定目标目录，就使用用户目录（'~'）作为目标目录
  }
  fmt.Println("change to " + destDirectory)
  return destDirectory, nil
}
