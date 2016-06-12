package command

import (
  "fmt"
)

var helpTable map[string]string = make(map[string]string)

func init() {

helpTable["help"] = `Usage: help [command]
command:
  cat | cd | cp | find | help | ls | mkdir | more | mv | quit | rm | version

Type 'quit' to exit the program.
`

helpTable["ls"] = `Usage: ls [-alht] [path]

options:
  -l
    Show in long format:
    [file-type] [permission] [create-time] [modify-time] [size] [author] [name]
  -a
    Show all files, including '.', '..' and hidden files.
  -h
    Show file size in human readable format.
  -t
    Sort result list order by modify time.

path:
  A string which is the path of the directory to be list. Path must exist.
  If path is not specific, current directory will be used.
`

helpTable["rm"] = `Usage: rm [-rfvdi] path

options:
  -r
    Recursive remove sub directories.
  -f
    Force remove, don't ask.
  -v
    Verbose, show what is going on.
  -d | --dry-run
    Just dry run, remove nothing, and show the expected result.
  -i
    Interactively ask when trying to remove an unempty directory.
    Answer 'y/yes' will do remove action, answer 'n/no' will do nothing.

path:
  A string which is the path of the directory/file to be removed. Path must exist.
`

helpTable["mkdir"] = `Usage: mkdir [-pv] path

options:
  -p
    Recursively create sub folder according to the path.
  -v
    Verbose, show what is happening.

path:
  A string which is the path of the directory to be created.
`

helpTable["cd"] = `Usage: cd [path]

This command has no options.

path:
  A string which is the path of the directory to change to.
`

helpTable["quit"] = `Usage: quit

This command has no options.
`

helpTable["cp"] = `Usage: cp [-nR] sources ... dest

options:
  -n
    Do not overwrite an existing file.
  -R
    Copy the directory and the entire subtree.

sources
  Source files.

dest
  Destination directory or file path.
`

}

func CommandHelp(options map[string]string, rest []string, cwd string) {
  if len(rest) == 0 {
    fmt.Print(helpTable["help"])
  } else {
    cmd := rest[0]
    if helpInformation, ok := helpTable[cmd]; ok {
      fmt.Print(helpInformation)
    } else {
      fmt.Print(helpTable["help"])
    }
  }
}
