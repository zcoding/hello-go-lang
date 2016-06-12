package command

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "../parse"
)

func checkCmdResult(err error) {
  if err != nil {
    fmt.Println(err.Error())
  }
}

func execCommand(cmd string, options map[string]string, rest []string, cwd string) string {
  var destPath string = cwd
  switch cmd {
    case "cd":
      _destPath, err := CommandCd(options, rest, cwd)
      destPath = _destPath
      checkCmdResult(err)
    case "cp":
      CommandCp(options, rest, cwd)
    case "ls":
      CommandList(options, rest, cwd)
    case "mkdir":
      CommandMkdir(options, rest, cwd)
    case "mv":
      CommandMove(options, rest, cwd)
    case "rm":
      CommandRemove(options, rest, cwd)
    case "help":
      CommandHelp(options, rest, cwd)
    case "version":
      CommandVersion(options, rest, cwd)
    case "quit":
      fmt.Println("The command 'quit' has no options.")
      CommandHelp(map[string]string{}, []string{}, cwd)
    default:
      fmt.Println("Command not found: ", cmd)
      CommandHelp(map[string]string{}, []string{}, cwd)
  }
  return destPath
}

func CommandLineStart() {
  reader := bufio.NewReader(os.Stdin)
  currentWorkingDirectory := "~"
  for {
    fmt.Print("> " + currentWorkingDirectory + " ")
    text, _ := reader.ReadString('\n')
    textTrim := strings.Trim(text, "\n")
    cmds := strings.Split(textTrim, " ")
    cmdsLen := len(cmds)
    if cmdsLen == 0 {
      continue
    }
    cmd := cmds[0]
    params := cmds[1:]
    if "quit" == cmd && cmdsLen == 1 {
      fmt.Println("Bye!")
      break
    }
    options, rest, err := parse.ParseOptions(params)
    if err != nil {
      fmt.Println(err.Error())
      execCommand("help", map[string]string{}, []string{cmd}, currentWorkingDirectory)
    } else {
      destPath := execCommand(cmd, options, rest, currentWorkingDirectory)
      currentWorkingDirectory = destPath
    }
  }
}
