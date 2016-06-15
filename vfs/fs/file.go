package fs

const BLOCKSIZE int = 2048 // 一个数据块的大小 2048B
const INDEXSIZE int = 9 // 索引数组的大小 9
const DENTRY_MAX int = 576 // 单个文件目录的目录项最大个数

/* 系统块 2048B */
type sys_block struct {
  system uint32               //文件系统标识      4B
  use_inode uint32            //已使用节点数      4B
  free_inode uint32           //空闲节点数        4B
  use_block uint32            //已使用数据块      4B
  free_block uint32           //空闲数据块       4B
  use_size uint32             //已使用空间       4B
  free_size uint32            //剩余空间        4B
  fill [2020]byte             //填充字节        2020B
}

/* 索引节点存储结构 64B */
type disk_inode struct {
  di_addr [INDEXSIZE]rune     //数据块地址 4B x 9 = 36B
  di_type uint16              //文件类型/目录标识 2B
  di_mood uint16              //父目录节点 2B
  di_add uint32             //二级索引标记 4B
  di_size uint32              //文件字节数 4B
  di_ctime uint64            //创建时间 8B
  di_mtime uint64            //修改时间 8B
}

/* 索引节点 */
type inode struct {
  i_node disk_inode         //磁盘i节点数据
  id_inode uint32               //磁盘i节点号
}

/* 目录项存储结构 32B */
type dentry struct {
  d_name [12]byte            //文件名 12B
  d_dir [12]byte             //所在目录 12B
  d_file uint32             //文件/目录标识 4B
  d_inode uint32              //磁盘i节点号 4B
}

/* 目录存储结构 1024B */
type dir struct {
  dentry [DENTRY_MAX]dentry   //当前目录下的所有文件目录项
  count uint16              //当前目录下的子目录和文件个数 2B
}
