/*
 * @Date: 2020-11-08 19:04:14
 * @Author: fenggq
 * @LastEditors: fenggq
 * @LastEditTime: 2020-11-11 20:04:52
 * @FilePath: /godemo/main.go
 */
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var (
	//AppName 应用名称
	AppName string
	//AppVersion 应用版本
	AppVersion string
	//BuildVersion 编译版本
	BuildVersion string
	//BuildTime 编译时间
	BuildTime string
	//GitRevision Git版本
	GitRevision string
	//GitBranch Git分支
	GitBranch string
	//GoVersion Golang信息
	GoVersion string
)

//GoTest ...
type GoTest struct {
	Responser      string   `json:"responser"`
	Version        string   `json:"version"`
	API            string   `json:"api"`
	Cmd            string   `json:"cmd"`
	ExpectedOutput []string `json:"expectedOutput"`
	Valid          bool     `json:"valid"`
}

//CMD ...
func CMD(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	data, err := cmd.Output()
	if err != nil {
		fmt.Printf("cmd:%s, get error %s\n", command, err.Error())
	}
	return (string)(data), err
}

//
func gotest(param *JenkinsMessageParam) string {
	out, err := CMD("go test")
	array := strings.Split(out, "\n")
	for _, v := range array {
		if strings.Contains(v, "responser") {
			testError := strings.SplitN(v, ": ", 2)
			for ke, ve := range testError {
				if strings.Contains(v, "responser") {
					msg := &GoTest{}
					err := json.Unmarshal([]byte(ve), msg)
					if err != nil {
						log.Println(err)
					}
					param.GoTestError = append(param.GoTestError, msg)
				}
				log.Println(ke, ve)
			}
		}
		//log.Println(k, v)
	}
	log.Println(err, "=====", out)
	return out
}

func main() {
	var token string
	log.SetFlags(log.Lshortfile)
	param := &JenkinsMessageParam{}
	testErr := gotest(param)
	flag.StringVar(&token, "t", "07de3a3799f70778bb98f95e7ef64b1693b30415a3ec59ea42e97de873f1aee0", "用户名,默认为空")
	flag.Parse()
	dingdingHook := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", token)
	log.Println(dingdingHook)
	dingTalk := DingTalk{
		Robot: Robot{
			WebHook: dingdingHook,
		},
	}
	user := LoadLatestCommitUser()
	param.GitCommitName = user
	param.ErrorMsg = testErr
	dingTalk.SendJenkinsMessage(param)
}
