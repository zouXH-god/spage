package models

import (
	"gorm.io/gorm"
	"time"
)

type JsonWebToken struct {
	gorm.Model
	UserID uint
}

type Permission struct {
	gorm.Model
	TargetType string
	TargetID   uint
}

type ApiToken struct {
	gorm.Model
	UserID    uint
	Token     string
	ExpiresAt time.Time
}
