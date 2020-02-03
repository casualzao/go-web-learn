package response

type LoginResponse struct {
	Token string `json:"token"`
	Uid   int    `json:"uid"`
}
