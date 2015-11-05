package types

type Database struct {
	Dbtype   string
	Server   string
	Dbname   string
	ConnMax  int
	User     string
	Password string
}

type AlipayConfig struct {
	Partner   string // 合作者ID
	Key       string // 合作者私钥
	ReturnUrl string // 同步返回地址
	NotifyUrl string // 网站异步返回地址
	Email     string // 网站卖家邮箱地址
}

type ServiceConfig struct {
	Port string //服务端口
}

type EnvConfig struct {
	DB      Database      `toml:"database"`
	Alipay  AlipayConfig  `tomal:"alipay"`
	SConfig ServiceConfig `tomal:"server"`
}
