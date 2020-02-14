package model

type CheckUserInfo struct {
	UserName           string `json:"user_name"`
	Email              string `json:"email"`
	ResultUserNameNum  int    `json:"result_user_name_num"`
	ResultEmailNum     int    `json:"result_email_num"`
	ResultUserNameText string `json:"result_user_name_text"`
	ResultEmailText    string `json:"result_email_text"`
}
