package types

type Result struct {
	// 状态
	Status int
	// 本网站订单号
	OrderNo string
	// 支付宝交易号
	TradeNo string
	// 买家支付宝账号
	BuyerEmail string
	// 错误提示
	Message string
}

// 生成订单的参数
type Options struct {
	OrderId  string  // 订单唯一id
	Fee      float32 // 价格
	NickName string  // 充值账户名称
	Subject  string  // 充值描述
}

type AlipayParametersWeb struct {
	InputCharset string  `json:"_input_charset"` //网站编码
	Body         string  `json:"body"`           //订单描述
	NotifyUrl    string  `json:"notify_url"`     //异步通知页面
	OutTradeNo   string  `json:"out_trade_no"`   //订单唯一id
	Partner      string  `json:"partner"`        //合作者身份ID
	PaymentType  uint8   `json:"payment_type"`   //支付类型 1：商品购买
	ReturnUrl    string  `json:"return_url"`     //回调url
	SellerEmail  string  `json:"seller_email"`   //卖家支付宝邮箱
	Service      string  `json:"service"`        //接口名称
	Subject      string  `json:"subject"`        //商品名称
	TotalFee     float32 `json:"total_fee"`      //总价
	Sign         string  `json:"sign"`           //签名，生成签名时忽略
	SignType     string  `json:"sign_type"`      //签名类型，生成签名时忽略
}

type AlipayParametersWap struct {
	InputCharset string  `json:"_input_charset"` //网站编码
	Body         string  `json:"body"`           //订单描述
	NotifyUrl    string  `json:"notify_url"`     //异步通知页面
	OutTradeNo   string  `json:"out_trade_no"`   //订单唯一id
	Partner      string  `json:"partner"`        //合作者身份ID
	PaymentType  uint8   `json:"payment_type"`   //支付类型 1：商品购买
	ReturnUrl    string  `json:"return_url"`     //回调url
	SellerId     string  `json:"seller_id"`      //卖家支付宝唯一用户号
	Service      string  `json:"service"`        //接口名称
	Subject      string  `json:"subject"`        //商品名称
	TotalFee     float32 `json:"total_fee"`      //总价
	Sign         string  `json:"sign"`           //签名，生成签名时忽略
	SignType     string  `json:"sign_type"`      //签名类型，生成签名时忽略
}

type AlipayParametersWebBank struct {
	InputCharset string  `json:"_input_charset"` //网站编码
	Defaultbank  string  `json:"defaultbank"`    //默认网银
	NotifyUrl    string  `json:"notify_url"`     //异步通知页面
	OutTradeNo   string  `json:"out_trade_no"`   //订单唯一id
	Partner      string  `json:"partner"`        //合作者身份ID
	PaymentType  uint8   `json:"payment_type"`   //支付类型 1：商品购买
	Paymethod    string  `json:"paymethod"`      //默认支付方式
	ReturnUrl    string  `json:"return_url"`     //回调url
	SellerEmail  string  `json:"seller_email"`   //卖家支付宝邮箱
	Service      string  `json:"service"`        //接口名称
	Subject      string  `json:"subject"`        //商品名称
	TotalFee     float32 `json:"total_fee"`      //总价
	Sign         string  `json:"sign"`           //签名，生成签名时忽略
	SignType     string  `json:"sign_type"`      //签名类型，生成签名时忽略
}