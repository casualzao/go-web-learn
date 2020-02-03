package common

import (
	"fmt"
	"go-web-learn/model"
	"log"
	"reflect"
	"testing"
)

func TestProcess(t *testing.T) {

	fmt.Println(*process)

}

func adder() func(int) int {
	sum := 0
	innerfunc := func(x int) int {
		sum += x
		return sum
	}
	return innerfunc
}

func Refac(fun interface{}, paramsValue interface{}) {

	values := getValues(paramsValue)
	v := reflect.ValueOf(fun)
	if v.Kind() != reflect.Func {
		log.Fatal("funcInter is not func")
	}

	val := v.Call(values)

	for i := range val {
		fmt.Println(val[i])
	}

}

//根据参数获取对应的Values

func G() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
		fmt.Println("c")
	}()
	F()
	fmt.Println("继续执行")
}

func F() {
	//defer func() {
	//	fmt.Println("b")
	//}()
	panic("a")
}

func TestAdder(t *testing.T) {

	//Refac(Add,model.User{"aa", "11"})

	G()

}
