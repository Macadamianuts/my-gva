package common

import (
	"gorm.io/gen"
	"gorm.io/gorm"
)

type (
	GenScopes  func(tx gen.Dao) gen.Dao
	GormScopes func(tx *gorm.DB) *gorm.DB
)
