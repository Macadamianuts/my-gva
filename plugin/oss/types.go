package oss

type OssType int

// 定义oss存储类型值
const (
	Minio OssType = iota
	AliyunOss
	AwsS3Oss
	HuaWeiObs
	QiniuKodo
	TencentCos
	LocalStore
)
