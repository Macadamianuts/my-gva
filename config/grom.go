package config

import "time"

// Dsn 默认走mysql
func (x *Gorm) Dsn() string {
	switch x.Type {
	case "mysql":
		return x.MysqlDsn()
	case "oracle":
		return x.OracleDsn()
	}
	return x.MysqlDsn()
}

// IsEmpty 返回数据是否为空, 为是否初始化数据做准备
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (x *Gorm) IsEmpty() bool {
	return x.Dbname == ""
}

// GetMaxIdleConesInt 获取 Gorm.MaxIdleCones int类型
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (x *Gorm) GetMaxIdleConesInt() int {
	return int(x.MaxIdleCones)
}

// GetMaxOpenConesInt 获取 Gorm.MaxOpenCones int类型
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (x *Gorm) GetMaxOpenConesInt() int {
	return int(x.MaxOpenCones)
}

// GetConnMaxLifetimeDuration 获取 Gorm.ConnMaxLifetime int类型
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (x *Gorm) GetConnMaxLifetimeDuration() time.Duration {
	return time.Duration(x.ConnMaxLifetime)
}

// GetConnMaxIdleTimeDuration 获取 Gorm.ConnMaxIdleTime int类型
//
//	Author [SliverHorn](https://github.com/SliverHorn)
func (x *Gorm) GetConnMaxIdleTimeDuration() time.Duration {
	return time.Duration(x.ConnMaxIdleTime)
}
