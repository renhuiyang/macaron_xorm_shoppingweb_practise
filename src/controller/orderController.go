package controller

import (
	"github.com/Unknwon/macaron"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	. "model/types"
	"net/http"
	"net/url"
	"strconv"
	"service/alipay"
	"encoding/json"
)

func PostOrder(ctx *macaron.Context, x *xorm.Engine, o Order) {
	alipaytype := ctx.Params("alipaytype")
	
	_, err := x.Insert(o)
	if err != nil {
		glog.V(1).Infof("Insert order %#v fail:%s", o, err.Error())
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	
	//after saved to db,we call alipay and submit to alipay
	outhtml, outscript := alipay.Form(o,alipaytype)

	ob := map[string]string{"html": outhtml, "script": outscript}
	ctx.Resp.Header().Set("Content-Type", "application/json")
	js, _ := json.Marshal(ob)

	ctx.Resp.WriteHeader(http.StatusOK)
	ctx.Resp.Write([]byte(js))
	return
}

func GetOrders(ctx *macaron.Context, x *xorm.Engine) {
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

	os := make([]Order, 0)
	err = x.Limit(limit, skip).Find(&os)
	if err != nil {
		glog.V(1).Infof("Get order from db fail:%s", err.Error())
		ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, os)
}