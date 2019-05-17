package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Base struct {
	gorm.Model
	Token        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
}
