package main

import (
	"github.com/BurntSushi/toml"
	"github.com/Unknwon/macaron"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"github.com/macaron-contrib/binding"
	//"github.com/macaron-contrib/cache"
	//_ "github.com/macaron-contrib/cache/redis"
	"github.com/macaron-contrib/session"
	_ "github.com/macaron-contrib/session/redis"

	"html/template"

	"flag"
	"fmt"
	"time"

	"controller"
	c "model/config"
	. "model/types"
)

var x *xorm.Engine

func initDB(config EnvConfig) {
	var err error
	mysqlserver := fmt.Sprintf("%v:%v@%v/%v?charset=utf8&parseTime=true", config.DB.User, config.DB.Password, config.DB.Server, config.DB.Dbname)
	x, err = xorm.NewEngine("mysql", mysqlserver)
	if err != nil {
		glog.V(1).Infof("Connect DataBase Fail:%v", err.Error())
	}

	x.ShowSQL = true //打印SQL语句

	//创建表
	err = x.Sync2(new(Customer))
	if err != nil {
		glog.V(1).Infof("Sync DataBase Fail:%v", err.Error())
	}

	err = x.Sync2(new(Goods))
	if err != nil {
		glog.V(1).Infof("Sync DataBase Fail:%v", err.Error())
	}

	err = x.Sync2(new(Order))
	if err != nil {
		glog.V(1).Infof("Sync DataBase Fail:%v", err.Error())
	}
}

func initParse() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
	flag.Lookup("log_dir").Value.Set("/tmp/log")
	flag.Lookup("v").Value.Set("4")
}

func readConfig(path string, config *EnvConfig) error {
	if _, err := toml.DecodeFile(path, config); err != nil {
		glog.V(1).Infof("[DEBUG] read configure file fail:%v!", err.Error())
		return err
	}
	return nil
}

func main() {
	initParse()

	err := readConfig("service.conf", &c.Config)
	if err != nil {
		glog.V(1).Infof("Read Configure file fail:%#v", err)
		return
	}

	initDB(c.Config)

	m := macaron.Classic()
	macaron.Env = macaron.PROD
	m.Use(macaron.Renderer())
	m.Use(session.Sessioner(session.Options{
		Provider:       "redis",
		ProviderConfig: "addr=127.0.0.1:6379",
	}))

	m.Use(macaron.Static("public",
		macaron.StaticOptions{
			// 请求静态资源时的 URL 前缀，默认没有前缀
			Prefix: "public",
			// 禁止记录静态资源路由日志，默认为不禁止记录
			SkipLogging: true,
			// 当请求目录时的默认索引文件，默认为 "index.html"
			IndexFile: "index.html",
			// 用于返回自定义过期响应头，不能为不设置
			// https://developers.google.com/speed/docs/insights/LeverageBrowserCaching
			Expires: func() string {
				return time.Now().Add(24 * 60 * time.Minute).UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
			},
		}))

	m.Use(macaron.Renderer(macaron.RenderOptions{
		// 模板文件目录，默认为 "templates"
		Directory: "templates",
		// 模板文件后缀，默认为 [".tmpl", ".html"]
		Extensions: []string{".tmpl", ".html"},
		// 模板函数，默认为 []
		Funcs: []template.FuncMap{map[string]interface{}{
			"AppName": func() string {
				return "Macaron"
			},
			"AppVer": func() string {
				return "1.0.0"
			},
		}},
		// 模板语法分隔符，默认为 ["{{", "}}"]
		Delims: macaron.Delims{"{{", "}}"},
		// 追加的 Content-Type 头信息，默认为 "UTF-8"
		Charset: "UTF-8",
		// 渲染具有缩进格式的 JSON，默认为不缩进
		IndentJSON: true,
		// 渲染具有缩进格式的 XML，默认为不缩进
		IndentXML: true,
		// 渲染具有前缀的 JSON，默认为无前缀
		PrefixJSON: []byte(""),
		// 渲染具有前缀的 XML，默认为无前缀
		PrefixXML: []byte(""),
		// 允许输出格式为 XHTML 而不是 HTML，默认为 "text/html"
		HTMLContentType: "text/html",
	}))
	m.Map(x)
	m.SetDefaultCookieSecret("UYCNJSHA123JCN409")
//	m.Post("/login",controller.Login)
//	m.Get("/logout",controller.Logout)

	m.Group("/users", func() {
		m.Post("", binding.Bind(Customer{}), controller.PostCustomer)
		m.Get("", controller.GetCustomers)
	})

	m.Group("/goods", func() {
		m.Post("", binding.Bind(Goods{}), controller.PostGood)
		m.Get("", controller.GetGoods)
	})

	m.Group("/orders", func() {
		m.Post("/alipay", binding.Bind(Order{}), controller.PostOrder)
		m.Get("", controller.GetOrders)
	})

	m.Group("/alipay", func() {
		m.Post("/notify", controller.AlipayNotify)
		m.Get("/return", controller.AlipayReturn)
	})

	m.Run(8080)
}
