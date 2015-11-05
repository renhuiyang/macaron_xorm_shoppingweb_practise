package types

import (
	"time"
)

type Order struct {
	Id         int64
	GoodId     int64      `json:"goodId" binding:"Required" xorm:"notnull"`
	GoodCnt    int64      `json:"goodcount" binding:"Required" xorm:"notnull"`
	Desc       string     `json:"desc" xorm:"varchar(255)"`
	CusId      int64      `json:"cusId"`
	CusName    string     `json:"cusname" binding:"Required" xorm:"notnull"`
	CusTel     string     `json:"custel" binding:"Required" xorm:"notnull"`
	CusAddr    string     `json:"cusaddr" binding:"Required" xorm:"notnull"`
	Status     string      `xorm:"notnull"`
	CreateTime time.Time   `xorm:"created"`
	PayTime    time.Time   `xorm:"updated"`
	Sum        float32     `xorm:"notnull"`
	Uuid       string      `xorm:"notnull unique"`
	ShippingID string      
	ShippingCom string     
}


