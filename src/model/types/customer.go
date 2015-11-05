package types

import (
	"time"
)

type Customer struct {
	Id       int64
	Name     string    `json:"name" binding:"Required" xorm:"notnull unique"`
	Telphone string    `json:"tel"  binding:"Required;MinSize(11)" xorm:"notnull unique"`
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
	PassWord string    `json:"pw" binding:"Required;MinSize(6)" xorm:notnull`
	Email    string    `json:"email" binding:"Required;Email"`
	Address  string    `json:"addr" xorn:"varchar(255)"`
	Version  int       `xorm:"version"`
}


