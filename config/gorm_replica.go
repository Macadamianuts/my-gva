package config

func (x *GormReplica) DataInterfaces() []any {
	length := len(x.Data)
	data := make([]any, 0, length)
	for i := 0; i < length; i++ {
		data = append(data, x.Data[i])
	}
	return data
}
