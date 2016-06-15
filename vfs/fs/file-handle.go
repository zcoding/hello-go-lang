package fs

/**
 * 按字节写
 * @param {int} index 写开始位置
 * @param {[]byte} block 写入的字节
 * @return {error} 可能发生错误
 */
func zfsWrite(index int, block []byte) error {
  return nil
}

/**
 * 按字节读
 * @param {int} index 读开始位置
 * @param {int} size 读取字节数
 * @return {[]byte} 读取的字节
 * @return {error} 可能发生错误
 */
func zfsRead(index int, size int) ([]byte, error) {
  return nil, nil
}
