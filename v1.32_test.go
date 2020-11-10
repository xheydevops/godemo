/*
 * @Date: 2020-11-10 12:01:54
 * @Author: fenggq
 * @LastEditors: fenggq
 * @LastEditTime: 2020-11-10 17:09:14
 * @FilePath: /godemo/v1.32_test.go
 */
package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

type caseItem struct {
	Responser      string   `json:"responser"`
	Version        string   `json:"version"`
	API            string   `json:"api"`
	Cmd            string   `json:"cmd"`
	ExpectedOutput []string `json:"expectedOutput"`
	Valid          bool     `json:"valid"`
}

var defaultResponser = "骆玉霞"
var casesV132 = []caseItem{
	{
		Responser:      "骆玉霞",
		Version:        "v1.32",
		API:            "/next/workgroup/v5/info/oneday",
		Cmd:            `curl https://test.xhey.top/next/workgroup/v5/info/oneday --data '{"userID":"xuser-c042091b-330d-494c-98a3-c95d726201de","groupID":"a06ee177-0e89-4ebe-8320-8929df0d127c","time":"2020-09-09","sign":"9d1d452deab1f16ac4deb4f411b04435","pageStartID":"","pageStartTime":""}'`,
		ExpectedOutput: []string{`20c9044d1a7592783fb2b777c21c032e.jpg`, `"mediaType":0`, `"sourceType":0`},
		Valid:          true,
	},
}

func cmd(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	data, err := cmd.Output()
	if err != nil {
		fmt.Printf("cmd:%s, get error %s\n", command, err.Error())
	}
	return (string)(data), err
}

func TestV132(t *testing.T) {
	for _, v := range casesV132 {
		caseItem, _ := json.Marshal(v)
		t.Errorf("%s", string(caseItem))
		return
		if v.Responser == "" {
			v.Responser = defaultResponser
		}
		data, err := cmd(v.Cmd)
		if err != nil {
			t.Errorf("%v", err)
		}
		for _, w := range v.ExpectedOutput {
			if strings.Contains(data, w) != true {
				caseItem, _ := json.Marshal(v)
				t.Errorf("%s", string(caseItem))
			}
		}
	}
}
