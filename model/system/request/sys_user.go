package request

type Register struct {
	Username string `json:"userName"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
}
