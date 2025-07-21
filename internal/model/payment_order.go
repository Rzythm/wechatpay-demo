package model

import (
	"gorm.io/gorm"
)

type OrderStatus string

const (
	StatusPending OrderStatus = "PENDING"
	StatusSuccess OrderStatus = "SUCCESS"
	StatusFailed  OrderStatus = "FAILED"
)

// PaymentOrder 存储用户的每一笔支付订单
type PaymentOrder struct {
	gorm.Model
	OutTradeNo  string      `gorm:"type:varchar(128);uniqueIndex;not null;comment:商户系统内部订单号"`
	OpenID      string      `gorm:"type:varchar(128);index;not null;comment:支付用户的OpenID"`
	CustomerID  string      `gorm:"type:varchar(64);index;not null;comment:供热公司的客户编号"`
	Amount      int64       `gorm:"not null;comment:支付金额，单位为分"`
	Description string      `gorm:"type:varchar(256);comment:订单描述"`
	Status      OrderStatus `gorm:"type:varchar(32);not null;comment:订单状态 (PENDING, SUCCESS, FAILED)"`
}
