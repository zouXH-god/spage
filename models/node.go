package models

import "gorm.io/gorm"

type Node struct {
	gorm.Model
	Token string `gorm:"unique;not null"` // 节点创建Token
	Name  string `gorm:"not null"`        // 节点名称
	Host  string `gorm:"not null"`
	Port  int    `gorm:"not null"`
}
