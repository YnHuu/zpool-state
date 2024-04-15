package models

type Wechat struct {
	APPID      string `json:"appid,omitempty"`
	APPSecret  string `json:"appsecret,omitempty"`
	TemplateID string `json:"template_id,omitempty"`
	ToUser     string `json:"to_user,omitempty"`
}
