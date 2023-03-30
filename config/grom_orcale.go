package config

import "fmt"

func (x *Gorm) OracleDsn() string {
	return fmt.Sprintf(`oracle://%s:%s@%s:%d/%s?%s`, x.Username, x.Password, x.Host, x.Port, x.Dbname, x.Config)
}
