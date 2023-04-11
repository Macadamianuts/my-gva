package oss

import (
	"context"
	"io"
	"mime/multipart"
	"os"
)

// Oss 接口
type Oss interface {
	// DeleteFile 根据key删除文件
	DeleteFile(ctx context.Context, filename string) error
	// UploadByFile 通过 *os.File 上传文件到oss
	UploadByFile(ctx context.Context, file *os.File) (filepath string, filename string, err error)
	// UploadByHeader 通过 *multipart.FileHeader 上传文件到oss
	UploadByHeader(ctx context.Context, header *multipart.FileHeader) (filepath string, filename string, err error)
}

// Plus 接口
type Plus interface {
	Oss
	// Upload 通过 io.Reader 上传文件到oss
	Upload(ctx context.Context, name string, reader io.Reader) (filepath string, filename string, err error)
}

// NewOss Oss 的实例化方法
func NewOss(ot OssType) Oss {
	switch ot {
	case Minio:
		return MinIO
	case AliyunOss:
		return Aliyun
	case QiniuKodo:
		return Qiniu
	case TencentCos:
		return Tencent
	case LocalStore:
		return Local
	default:
		return Local
	}
}
