package controller

import (
	"github.com/Unknwon/macaron"
	"github.com/go-xorm/xorm"
	. "model/types"
	"net/url"
	"service/alipay"
	"strconv"
)

//process post alipay return
func AlipayReturn(ctx *macaron.Context, x *xorm.Engine) {
	m, _ := url.ParseQuery(ctx.Req.URL.RawQuery)
	params := map[string]string{}
	for k, v := range m {
		params[k] = v[0]
	}
	result := alipay.Return(params)

	type OrderInfo struct {
		Result  bool
		OrderId string
		GoodId  string
		GoodCnt int64
		Tel     string
		Name    string
		Addr    string
		Sum     float32
	}

	var orderInfo OrderInfo
	orderInfo.Result = false
	if result.Status == -1 || result.Status == -5 || result.Status == -3 {
		ctx.HTML(400, "orderresult", orderInfo)
		return
	}

	o := &Order{Uuid: result.OrderNo}
	has, err := x.Where("status=?","等待支付").Get(o)
	if err != nil || !has {
		ctx.HTML(400, "orderresult", orderInfo)
		return
	}

	if result.Status != 1 {
		o.Status = "支付失败"
		x.Id(o.Id).Cols("status").Update(o)
		ctx.HTML(400, "orderresult", orderInfo)
		return
	}

	o.Status = "支付成功"
	_,err = x.Id(o.Id).Cols("status").Update(o)
	if err != nil {
		ctx.HTML(400, "orderresult", orderInfo)
		return
	}

	orderInfo.Result = true
	orderInfo.OrderId = o.CusTel + "_" + strconv.FormatInt(o.Id, 10)
	orderInfo.GoodId = strconv.FormatInt(o.GoodId, 10)
	orderInfo.GoodCnt = o.GoodCnt
	orderInfo.Tel = o.CusTel
	orderInfo.Name = o.CusName
	orderInfo.Addr = o.CusAddr
	orderInfo.Sum = o.Sum

	ctx.HTML(200, "orderresult", orderInfo)
	return
}

//process post alipay notify
func AlipayNotify(ctx *macaron.Context, x *xorm.Engine) {
	// Read the content
	bodyString, err := ctx.Req.Body().String()
	if err != nil {
		ctx.Resp.Write([]byte("error"))
		return
	}

	result := alipay.Notify(bodyString)

	if result.Status == -2 {
		ctx.Resp.Write([]byte("error"))
		return
	}
	
	o := &Order{Uuid: result.OrderNo}
	has, err := x.Where("status=?","等待支付").Get(o)
	if err != nil || !has {
		ctx.Resp.Write([]byte("error"))
		return
	}

	if result.Status != 1 {
		o.Status = "支付失败"
		x.Id(o.Id).Cols("status").Update(o)
        ctx.Resp.Write([]byte("error"))
		return
	}

	o.Status = "支付成功"
	_,err = x.Id(o.Id).Cols("status").Update(o)
	if err != nil {
		ctx.Resp.Write([]byte("error"))
		return
	}
	
	ctx.Resp.Write([]byte("success"))
	return
}
