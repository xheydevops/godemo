/*
 * @Date: 2020-11-10 16:46:16
 * @Author: fenggq
 * @LastEditors: fenggq
 * @LastEditTime: 2020-11-11 18:46:43
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
	GitUser = make(map[string]string)
	GitUser["andyfenggq"] = "冯国庆"
	GitUser["Aleutian Xie"] = "谢辉生"
	GitUser["luo"] = "骆玉霞"
	GitUser["audu"] = "杜于庆"
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

//SendJenkinsMessage ...
func (d *DingTalk) SendJenkinsMessage(param *JenkinsMessageParam, result *GoTest) (WebHookResponse, error) {
	title := "jenkins自动化测试"
	user := GitUser[param.GitCommitName]
	commitUser := MericsUser[user]
	responser := d.GetReporters(param)

	text := fmt.Sprintf("### 最近提交者：%s \n 开发者：%s \n ### gitlog ```json\n %s \n``` ### 测试错误信息\n```json\n %s \n``` \n @%s @%s",
		user, result.Responser, param.GitLog, param.ErrorMsg, commitUser, responser)

	msg := MarkdownMessage{
		Title: title,
		Text:  text,
	}
	at := At{
		AtMobiles: []string{"17316225231",
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
	return d.send(body)
}
