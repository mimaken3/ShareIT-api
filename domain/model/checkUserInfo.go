package model

type CheckUserInfo struct {
	UserName           string `json:"userName"`
	Email              string `json:"email"`
	ResultUserNameNum  int    `json:"resultUserNameNum"`
	ResultEmailNum     int    `json:"resultEmailNum"`
	ResultUserNameText string `json:"resultUserNameText"`
	ResultEmailText    string `json:"resultEmailText"`
}
