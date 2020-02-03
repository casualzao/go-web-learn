package exception

var SUCCESS = New("0", "SUCCESS")
var ERROR = New("999999", "ERROR")
var USER_NOT_EXIST = New("100001", "用户不存在")
var USER_NOT_LOGIN = New("100002", "用户未登陆")

type Exception struct {
	ErrorCode string
	ErrorMsg  string
}

func New(errorCode string, errorMsg string) Exception {
	return Exception{errorCode, errorMsg}
}

//func (e *Exception) Error() string{
//
//	return fmt.Sprintf("【errorCode】'%q' ,【msg】'%q' ",e.ErrorCode,e.ErrorMsg)
//}
