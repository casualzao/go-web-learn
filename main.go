package main

import (
	"fmt"
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris/v12"
	gf "github.com/snowlyg/gotransformer"
	"go-web-learn/config"
	"go-web-learn/middleware"
	"go-web-learn/routes"
	"time"
)

var Sc iris.Configuration

func main() {
	Sc = iris.TOML("./config/conf.toml")
	conf := getSysConf()
	app := NewApp(conf)

	err := app.Run(iris.Addr(conf.App.Port), iris.WithConfiguration(Sc))

	if err != nil {
		fmt.Println(err)
	}
}

func NewApp(rc *config.Conf) *iris.Application {
	app := iris.New()
	app.Logger().SetLevel(rc.App.LoggerLevel)

	app.RegisterView(iris.HTML("resources", ".html"))
	app.HandleDir("/static", "resources/static")

	// 生称api 文档
	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       true,
		DocTitle: rc.App.Name,
		DocPath:  "./resources/apiDoc/index.html", //设置绝对路径
		BaseUrls: map[string]string{
			"Production": rc.App.URl + rc.App.Port,
			"Staging":    "",
		},
	})
	// 注册路由
	routes.Register(app)
	middleware.Register(app)
	getRoutes(app)
	return app
}

// 获取路由信息
func getRoutes(api *iris.Application) {
	rs := api.APIBuilder.GetRoutes()

	fmt.Println(rs)
}

func getSysConf() *config.Conf {

	// 初始化 app 信息
	app := config.App{}
	g := gf.NewTransform(&app, Sc.Other["App"], time.RFC3339)
	_ = g.Transformer()

	mysql := config.Mysql{}
	g.OutputObj = &mysql
	g.InsertObj = Sc.Other["Mysql"]
	_ = g.Transformer()

	redis := config.Redis{}
	g.OutputObj = &redis
	g.InsertObj = Sc.Other["redis"]
	_ = g.Transformer()

	cf := &config.Conf{
		App:   app,
		Mysql: mysql,
		Redis: redis,
	}
	return cf
}
