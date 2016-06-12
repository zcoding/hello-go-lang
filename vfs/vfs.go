package main

import (
  "./command"
  "flag"
  "fmt"
  "./fs"
  // "./parse"
)

const vfsSize int64 = 1024

func main() {
  path := flag.String("p", "", "The path of an existing zfs file.")
  newPath := flag.String("n", "", "The path of a new zfs file.")
  flag.Parse()
  if "" != *newPath {
    fs.TryCreateFS(*newPath, vfsSize)
  } else if "" != *path {
    fs.TryOpenFS(*path)
  } else {
    fmt.Println("Illegal parameters. You must specify the path of an existing vfs file with option '-p', or the path of a new vfs file with option '-n'.")
    return
  }
  command.CommandLineStart()
}
