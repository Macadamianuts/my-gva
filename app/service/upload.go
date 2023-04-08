package service

import (
	"context"
	"gva-lbx/plugin/oss"
	"mime/multipart"
)

func UploadFile(ctx context.Context, header *multipart.FileHeader) (path, fileName string, err error) {
	path, fileName, err = oss.NewOss(oss.LocalStore).UploadByHeader(ctx, header)
	if err != nil {
		return "", "", err
	}
	return path, fileName, nil
}
