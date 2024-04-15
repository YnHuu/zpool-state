package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"web/models"
)

func (s *service) GetWechat() models.Wechat {
	var wc models.Wechat
	f, err := os.ReadFile(s.exPath + "wechat.conf")
	if err != nil {
		log.Println(err)
		return wc
	}
	_ = json.Unmarshal(f, &wc)
	return wc
}

func (s *service) SetWechat(data models.Wechat) bool {
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return false
	}
	err = os.WriteFile(s.exPath+"wechat.conf", b, 0644)
	return err == nil
}

func (s *service) TestWechat() bool {
	return s.SendWechat("测试", "ok")
}

func (s *service) SendWechat(task, status string) bool {
	wechat := s.GetWechat()
	token, err := getToken(wechat.APPID, wechat.APPSecret)
	if err != nil {
		return false
	}
	return sendWechat(token, task, status, wechat)
}

func getToken(appid, appsecret string) (string, error) {
	uri := `https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=` + appid + `&secret=` + appsecret
	resp, err := http.Get(uri)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	var data map[string]any
	_ = json.Unmarshal(body, &data)
	return data["access_token"].(string), nil
}

func sendWechat(token, task, status string, wc models.Wechat) bool {
	uri := `https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=` + token
	msg := fmt.Sprintf(`{"touser": "%v",
    "template_id": "%v",
    "data": {"task":{"value":"%v"},"status":{"value":"%v"}}}`, wc.ToUser, wc.TemplateID, task, status)
	resp, err := http.Post(uri, "application/json", strings.NewReader(msg))
	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	var data map[string]any
	_ = json.Unmarshal(body, &data)
	return data["errmsg"].(string) == "ok"
}
