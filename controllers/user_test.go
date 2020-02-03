package controllers

import (
	"go-web-learn/common"
	"go-web-learn/exception"
	"go-web-learn/model"
	"go-web-learn/service"
	"testing"
)

func TestLogin(t *testing.T) {

}

func BenchmarkLogin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//conventer()
		common.Process(service.Login, model.User{"aa", "bb"})
	}
}

func conventer() common.Response {
	login := service.Login(model.User{"aa", "bb"})
	response := common.Response{login, exception.SUCCESS.ErrorCode, exception.SUCCESS.ErrorMsg}
	return response
}

/**
goos: darwin
goarch: amd64
pkg: go-web-learn/controllers
BenchmarkLogin-8   	61074237	        19.3 ns/op
PASS


goos: darwin
goarch: amd64
pkg: go-web-learn/controllers
BenchmarkLogin-8   	 2330212	       526 ns/op
PASS

Process finished with exit code 0
*/
