package main

import (
  "./command"
  "flag"
  "fmt"
  "./fs"
)

// const vfsSize int64 = 1024 * 1024 * 1024 // 1GB
const vfsSize int64 = 1024 // 1KB

func main() {
  path := flag.String("p", "", "The path of an existing vfs file.")
  newPath := flag.String("n", "", "The path of a new vfs file.")
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
