package common

import (
	"go-web-learn/exception"
	"log"
	"reflect"
)

/**
反射类型，错误异常封装处理
@param  fun 函数
*/
func Process(fun interface{}, param interface{}) (r *Response) {
	// 异常封装处理
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			// 系统定义异常
			case exception.Exception:
				of := reflect.TypeOf(err)
				valueOf := reflect.ValueOf(err)
				convert := valueOf.Convert(of)
				r = &Response{nil, convert.Field(0).String(), convert.Field(1).String()}
			// 其余异常
			default:
				log.Println(err)
				r = &Response{nil, exception.ERROR.ErrorCode, exception.ERROR.ErrorMsg}
			}
		}

	}()
	values := getValues(param)
	v := reflect.ValueOf(fun)

	//判断 方法是否一致
	if v.Kind() != reflect.Func {
		log.Fatal("funcInter is not func")
	}

	// 进行调用
	val := v.Call(values)
	// 对对西那个给值
	r = &Response{val[0].Interface(), exception.SUCCESS.ErrorCode, exception.SUCCESS.ErrorMsg}
	return
}

/**
interface{} 泛型 参数 转换 []reflect.Value 进行反射调用
*/
func getValues(param ...interface{}) []reflect.Value {
	values := make([]reflect.Value, 0, len(param))
	for i := range param {
		values = append(values, reflect.ValueOf(param[i]))
	}
	return values
}
