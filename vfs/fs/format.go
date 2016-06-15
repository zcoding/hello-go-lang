package fs

import (
  "fmt"
)

const BitmapStart int = 2048

/**
 * Format the zfs file system
 * 
 */
func FormatFS() error {
  fmt.Println("Fotmating...")
  // 初始化位图
  var bitmap [BLOCKSIZE]byte
  for i := 0; i < BLOCKSIZE; i++ {
    bitmap[i] = 0xFF
  }
  err := zfsWrite(BitmapStart, bitmap)
  // 节点位图从2048开始
  // 块位图从4096开始
  // 初始化索引节点
  // 初始化目录
  return nil
}
