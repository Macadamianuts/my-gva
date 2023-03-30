package oss

type OssType int

const (
	Minio OssType = iota
	AliyunOss
	AwsS3Oss
	HuaWeiObs
	QiniuKodo
	TencentCos
	LocalStore
)
