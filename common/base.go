package common

/**
通用返回 response

*/
type Response struct {
	Data      interface{} `json:"data"`
	ErrorCode string      `json:"errorCode"`
	ErrorMsg  string      `json:"errorMsg"`
}
