/*
 * @Date: 2020-11-10 16:40:52
 * @Author: fenggq
 * @LastEditors: fenggq
 * @LastEditTime: 2020-11-10 18:36:53
 * @FilePath: /godemo/robot.go
 */
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Robot 消息机器人结构体, 消息处理都由 Robot 完成
type Robot struct {
	WebHook string
}

const (
	pageSize             = 1000
	configKey            = "robot:enabled"
	photoNumberKeyFormat = "group:%s:user:%s:number:%s"
	refreshLocker        = "robot:refresh:locker"
	dingTalkRobot        = "dingtalk"
	wechatWorkRobot      = "wechatwork"
	dingTalkLimit        = 20
	wechatWorkLimit      = 20
	failedNotice         = "照片转发失败，打开今日水印相机App，在我的团队中可查看此照片。"
)

// TextMessage 机器人文本消息体
type TextMessage struct {
	Content string `json:"content"`
}

// MarkdownMessage 机器人 Markdown 消息体
type MarkdownMessage struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

//At ...
type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

// WebHookResponse 调用 web hook 之后返回的消息体
type WebHookResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// 发送消息到 Web hook
func (r *Robot) send(body map[string]interface{}) (WebHookResponse, error) {
	message, _ := json.Marshal(body)
	jsonResp := WebHookResponse{}
	url := strings.Replace(r.WebHook, "\n", "", -1)
	url = strings.TrimSpace(url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		return jsonResp, fmt.Errorf("Send message failed with error: %+v, web hook: %+v, message: %+v", err, r.WebHook, string(message))
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, respErr := client.Do(req)

	if respErr != nil {
		return jsonResp, fmt.Errorf("RobotSendMessageFailed, error: %+v", respErr)
	}

	if resp != nil {
		defer func() {
			if resp.Body != nil {
				_ = resp.Body.Close()
			}
		}()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		if resp.StatusCode != http.StatusOK {
			return jsonResp, fmt.Errorf("RobotSendMessageFailed, response: %+v", resp)
		}

		decodeErr := json.Unmarshal(bodyBytes, &jsonResp)

		if decodeErr != nil {
			return jsonResp, fmt.Errorf("RobotSendMessageFailed, body json unmarshal failed: %+v", decodeErr)
		}

		return jsonResp, nil
	}
	return jsonResp, errors.New("RobotSendMessageFailed, response is nil")
}
