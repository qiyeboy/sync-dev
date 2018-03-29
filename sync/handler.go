package sync

import "os"

// 处理器
type Handler interface {
	// 文件路径映射(lpath 本地路径)返回远端路径
	MapPath(lpath string) string
	// 在远端创建目录
	CreateDir(lpath string) error
	// 删除远端目录（递归）
	RemoveDir(lpath string) error
	// 删除远端文件
	RemoveFile(lpath string) error
	// 将本地文件传输到远端 lfile 为 lpath 对应的本地文件对象
	UploadFile(lpath string, lfile *os.File) error
	// 重命名
	Rename(opath, npath string) error
	// 关闭处理器
	Close() error
	// 目录完整同步（一般不清理远端（单独存在）的文件，例如在远端生成了日志文件）
	SyncDir(lpath string) error
}