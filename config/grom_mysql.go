package config

import "fmt"

// MysqlDsn mysql dsn
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (x *Gorm) MysqlDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", x.Username, x.Password, x.Host, x.Port, x.Dbname, x.MysqlConfig())
}

// MysqlConfig 获取 mysql 其他配置项
func (x *Gorm) MysqlConfig() string {
	if x.Config == "" {
		return "charset=utf8mb4&parseTime=True&loc=Local"
	}
	return x.Config
}

// MysqlEmptyDsn mysql 空dsn即不带数据库的dsn
func (x *Gorm) MysqlEmptyDsn() string {
	if x.Type != "mysql" {
		return x.PostgresEmptyDsn()
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)?%s", x.Username, x.Password, x.Host, x.Port, x.MysqlConfig())
}
