package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12/context"
)

func JwtHandLer() *jwt.Middleware {

	var mySecret = []byte("HS2JDFKhu8u1i76d")

	return jwt.New(jwt.Config{
		//jwt 校验key  配置获取信息
		ValidationKeyGetter: func(token *jwt.Token) (i interface{}, err error) {
			return mySecret, nil
		},
		// 上下文 key ？
		ContextKey: "",
		// 校验失败 handler
		ErrorHandler: func(context context.Context, err error) {

		},
		//todo ？
		CredentialsOptional: false,
		//todo ？
		Extractor: nil,
		//todo ？
		EnableAuthOnOptions: false,
		//签名方法
		SigningMethod: jwt.SigningMethodES256,
		//todo ？
		Expiration: false,
	})
}
