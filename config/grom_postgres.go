package config

import "fmt"

// PostgresDsn 获取 postgres dsn
func (x *Gorm) PostgresDsn() string {
	return fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s %s`, x.Host, x.Port, x.Username, x.Password, x.Dbname, x.PostgresConfig())
}

// PostgresConfig 获取 postgres 配置项
func (x *Gorm) PostgresConfig() string {
	if x.Config == "" {
		return "sslmode=disable TimeZone=Asia/Shanghai"
	}
	return x.Config
}

// PostgresEmptyDsn postgres 空dsn即不带数据库的dsn
func (x *Gorm) PostgresEmptyDsn() string {
	if x.Type != "postgres" {
		return x.MysqlEmptyDsn()
	}
	return fmt.Sprintf(`host=%s port=%d user=%s password=%s %s`, x.Host, x.Port, x.Username, x.Password, x.PostgresConfig())
}
