package controllers

import (
	"github.com/kataras/iris/v12"
	"go-web-learn/common"
	"go-web-learn/model"
	"go-web-learn/service"
)

/**
用户注册

*/
func Register(ctx iris.Context) {
	user := new(model.User)

	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		process := common.Process(service.Login, model.User{"aa", "bb"})

		_, _ = ctx.JSON(process)
	} else {
		//_, _ = ctx.JSON(common.ApiResponse("User", "Ok"))

	}
}

/**
用户登陆
*/
func Login(ctx iris.Context) {

}
