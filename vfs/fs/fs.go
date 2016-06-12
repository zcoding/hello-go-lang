package fs

import (
  "io/ioutil"
  "fmt"
  "os"
)

func checkFile(err error) {
  if err != nil {
    fmt.Println(err.Error())
    panic(err)
  }
}

func TryCreateFS(path string, size int64) {
  d1 := make([]byte, size)
  // 如果文件已经存在，抛出error
  err := ioutil.WriteFile(path, d1, 0644)
  checkFile(err)
}

func TryWriteBytes(path string, words string, offset int64) {
  zfs, err := os.OpenFile(path, os.O_RDWR, 0644)
  checkFile(err)
  defer zfs.Close()
  n, err := zfs.WriteAt([]byte(words), offset)
  checkFile(err)
  fmt.Println(n)
}

func TryOpenFS(path string) {
  zfs, err := os.OpenFile(path, os.O_RDWR, 0644)
  checkFile(err)
  defer zfs.Close()
}
