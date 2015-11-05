package types

import (
	"time"
)

type Goods struct {
	Id       int64
	Name     string    `json:"name" binding:"Required" xorm:"notnull unique"`
	Desc     string    `json:"desc" xorn:"varchar(255)"`
	Fee      float32   `json:"fee" binding:"Required" xorm:"notnull"` 
	Version  int       `xorm:"version"`
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
}


