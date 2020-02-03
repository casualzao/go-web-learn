package middleware

import "github.com/kataras/iris/v12/context"
import "github.com/iris-contrib/middleware/cors"

/**
请求权限设置
*/
func CrsAuth() context.Handler {
	return cors.New(cors.Options{
		// 跨域设置
		AllowedOrigins: []string{"*"},
		// 允许 请求的方法参数
		AllowedMethods: []string{"PUT", "GET", "POST", "OPTIONS"},
		// 允许的请求头
		AllowedHeaders: []string{"*"},
		// 报漏的请求头
		ExposedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		MaxAge:         0,
		//  允许证书
		AllowCredentials: true,
		// 跨域 配置，待定
		OptionsPassthrough: false,
		Debug:              false,
	})
}
