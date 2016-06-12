package main

import (
  "fmt"
  "image"
  "os"
)

func checkError(err error) {
  if err != nil {
    fmt.Println(err.Error())
  }
}

func main() {
  file1, err := os.Open("./test/lena.jpg")
  checkError(err)
  defer file1.Close()
  image1, _, err := image.Decode(file1)
  checkError(err)
}
