package controller

import (
	"github.com/Unknwon/macaron"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	. "model/types"
	"net/http"
	"net/url"
	"strconv"
)

func PostGood(ctx *macaron.Context, x *xorm.Engine, g Goods) {
	_, err := x.Insert(g)
	if err != nil {
		glog.V(1).Infof("Insert good %#v fail:%s", g, err.Error())
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, "SUCCESS")
	return
}

func GetGoods(ctx *macaron.Context, x *xorm.Engine) {
	m, _ := url.ParseQuery(ctx.Req.URL.RawQuery)
	glog.V(1).Infof("Debug %#v", m)
	skip := 0
	limit := 0
	var err error

	if v, ok := m["skip"]; ok {
		skip, _ = strconv.Atoi(v[0])
	}

	if v, ok := m["limit"]; ok {
		limit, _ = strconv.Atoi(v[0])
	}

	gs := make([]Goods, 0)
	err = x.Limit(limit, skip).Find(&gs)
	if err != nil {
		glog.V(1).Infof("Get good from db fail:%s", err.Error())
		ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gs)
}