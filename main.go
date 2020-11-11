/*
 * @Date: 2020-11-08 19:04:14
 * @Author: fenggq
 * @LastEditors: fenggq
 * @LastEditTime: 2020-11-11 14:43:54
 * @FilePath: /godemo/main.go
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
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

func main() {
	log.SetFlags(log.Lshortfile)

	//cmd := os.Args[0]
	name := os.Args[1]
	errMssage := ""
	for i, a := range os.Args[2:] {
		fmt.Printf("Argument %d is %s\n", i+1, a)
		if errMssage != "" {
			errMssage += " "
		}
		errMssage += a
	}
	log.Println(name, errMssage)

	msg := &GoTest{}
	err := json.Unmarshal([]byte(errMssage), msg)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%#v,%#v", msg.Responser, msg)
	dingTalk := DingTalk{
		Robot: Robot{
			WebHook: "https://oapi.dingtalk.com/robot/send?access_token=07de3a3799f70778bb98f95e7ef64b1693b30415a3ec59ea42e97de873f1aee0",
		},
	}
	gitlog := LoadCommitMessage("")
	//log.Println(gitlog)
	return
	//	out, err := json.Marshal(msg)
	param := &JenkinsMessageParam{}
	param.GitCommitName = name
	param.GitLog = gitlog
	param.ErrorMsg = errMssage
	dingTalk.SendJenkinsMessage(param, msg)
}
