package parse

import (
  "strings"
  "errors"
)

/**
 * 读取命令参数，解析成map和[]string
 * 命名参数解析成map，剩下的参数按顺序组成数组
 * @param {[]string} params 参数列表
 * @return {map[string]string} 命名参数表
 * @return {[]string} 剩余参数数组
 * @return {error} 遇到非法的参数时返回错误，否则返回nil
 */
func ParseOptions(params []string) (map[string]string, []string, error) {
  options := make(map[string]string)
  rest := []string{}
  for i := 0; i < len(params); i++ {
    if strings.HasPrefix(params[i], "-") {
      parts := strings.Split(params[i], "=")
      if "-" == parts[0] || "--" == parts[0] {
        return nil, nil, errors.New("Illegal option.")
      }
      if len(parts) == 2 {
        options[parts[0]] = parts[1]
      } else {
        options[parts[0]] = ""
      }
    } else {
      rest = append(rest, params[i])
    }
  }
  return options, rest, nil
}
