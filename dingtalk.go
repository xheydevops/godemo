/*
 * @Date: 2020-11-10 16:46:16
 * @Author: fenggq
 * @LastEditors: fenggq
 * @LastEditTime: 2020-11-12 10:54:14
 * @FilePath: /godemo/dingtalk.go
 */
package main

import (
	"fmt"
	"log"
	"strings"
)

//JenkinsMessageParam ...
type JenkinsMessageParam struct {
	GitCommitName string    `json:"commitName"`
	GoTestError   []*GoTest `json:"goTest"`
	GitLog        string    `json:"gitLog"`
	ErrorMsg      string    `json:"errorMsg"`
	Type          string    `json:"type"`
}

//MericsUser ...
var MericsUser map[string]string

//GitUser ...
var GitUser map[string]string

//gitEmailMap ...
var gitEmailMap map[string]string

// DingTalk ...
type DingTalk struct {
	Robot
}

func init() {
	gitEmailMap = make(map[string]string)
	gitEmailMap["362739259@qq.com"] = "冯国庆"
	gitEmailMap["aleutian.xie@cicisoft.cn"] = "谢辉生"
	gitEmailMap["luo_yu_xia@163.com"] = "骆玉霞"
	gitEmailMap["audu@qq.com"] = "杜于庆"
	gitEmailMap["icecut@qq.com"] = "杜于庆"
	GitUser = make(map[string]string)
	GitUser["andyfenggq"] = "冯国庆"
	GitUser["Aleutian Xie"] = "谢辉生"
	GitUser["luo"] = "骆玉霞"
	GitUser["audu"] = "杜于庆"
	GitUser["姜亦春"] = "姜亦春"
	MericsUser = make(map[string]string)
	MericsUser["冯国庆"] = "17316225231"
	MericsUser["骆玉霞"] = "13552079799"
	MericsUser["谢辉生"] = "15901435695"
	MericsUser["姜亦春"] = "13581894261"
	MericsUser["杜于庆"] = "18211025188"
}

//SendTextMessage ...
func (d *DingTalk) SendTextMessage(text string) (WebHookResponse, error) {
	msg := TextMessage{
		Content: "`" + text + "`",
	}
	body := make(map[string]interface{})
	body["msgtype"] = "text"
	body["text"] = msg

	return d.send(body)
}

//GetReporters ...
func (d *DingTalk) GetReporters(param *JenkinsMessageParam) string {
	var responserMap map[string]string
	responserMap = make(map[string]string)
	for _, v := range param.GoTestError {
		if len(v.Responser) > 0 {
			responserMap[v.Responser] = MericsUser[v.Responser]
		}
	}
	var Responsers []string
	for _, responserPhone := range responserMap {
		Responsers = append(Responsers, responserPhone)
	}
	responser := strings.Join(Responsers, "@")
	log.Println(responser)
	return responser
}

//SendMarkdownMessage ...
func (d *DingTalk) SendMarkdownMessage(param *JenkinsMessageParam) (WebHookResponse, error) {
	return WebHookResponse{}, nil
}

//SendJenkinsMessage ...
func (d *DingTalk) SendJenkinsMessage(param *JenkinsMessageParam) (WebHookResponse, error) {
	title := "jenkins自动化测试"
	log.Printf("%#v", param)
	user := GitUser[param.GitCommitName]
	commitUser := MericsUser[user]
	responser := d.GetReporters(param)
	alertUser := ""
	resultTitle := "测试错误信息"
	if len(responser) > 0 {
		if len(commitUser) > 0 && strings.Contains(responser, commitUser) == false {
			alertUser = fmt.Sprintf("@%s", commitUser)
		}
		alertUser = fmt.Sprintf("%s@%s", alertUser, responser)
	} else {
		resultTitle = "测试结果正常"
	}
	param.GitLog = LoadLatestGitLogs()
	text := fmt.Sprintf("### 最近提交者：%s \n \n ### gitlog \n```json\n %s \n```\n ### %s\n```json\n %s \n``` \n %s",
		user, param.GitLog, resultTitle, param.ErrorMsg, alertUser)

	msg := MarkdownMessage{
		Title: title,
		Text:  text,
	}
	at := At{
		AtMobiles: []string{
			"17316225231",
			"13552079799",
			"15901435695",
			"13581894261",
			"18211025188"},
		IsAtAll: false,
	}
	body := make(map[string]interface{})
	body["msgtype"] = "markdown"
	body["markdown"] = msg
	body["at"] = at
	responce, err := d.send(body)
	if err == nil && len(alertUser) > 0 {
		text := fmt.Sprintf("自动化build错误，请关注 %s", alertUser)
		d.SendTextMessage(text)
	}
	return responce, err
}
