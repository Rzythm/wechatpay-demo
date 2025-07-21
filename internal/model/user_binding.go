package model

import "gorm.io/gorm"

// UserBinding 存储微信用户OpenID与供热客户ID的绑定关系
type UserBinding struct {
	gorm.Model        // 包含了 ID, CreatedAt, UpdatedAt, DeletedAt 四个字段
	OpenID     string `gorm:"type:varchar(128);uniqueIndex;not null;comment:微信用户的OpenID"`
	CustomerID string `gorm:"type:varchar(64);not null;comment:供热公司的客户编号"`
}
