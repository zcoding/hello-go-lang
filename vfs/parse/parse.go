package parse

import (
  "strings"
  "errors"
)

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
