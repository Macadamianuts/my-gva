package config

import "fmt"

// MysqlDsn 获取 mysql dsn
func (x *GormReplica) MysqlDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", x.Username, x.Password, x.Host, x.Port, x.Dbname, x.MysqlConfig())
}

// MysqlConfig 获取 mysql 配置项
func (x *GormReplica) MysqlConfig() string {
	if x.Config == "" {
		return "charset=utf8mb4&parseTime=True&loc=Local"
	}
	return x.Config
}
