package config

import "fmt"

func (x *System) Address() string {
	return fmt.Sprintf(":%d", x.Port)
}
