/*
 * @Date: 2020-11-09 11:18:42
 * @Author: fenggq
 * @LastEditors: fenggq
 * @LastEditTime: 2020-11-30 10:49:32
 * @FilePath: /godemo/go_test.go
 */
package main

import (
	"testing"
)

//TestRun ...
func TestRun(t *testing.T) {
	/*	log.Printf("App Name:\t%s\n", AppName)
		log.Printf("App Version:\t%s\n", AppVersion)
		log.Printf("Build version:\t%s\n", BuildVersion)
		log.Printf("Build time:\t%s\n", BuildTime)
		log.Printf("Git revision:\t%s\n", GitRevision)
		//log.Printf("Git branch:\t%s\n", GitBranch)
		log.Printf("Golang Version: %s\n", GoVersion)*/
	t.Log("test", "TestRun")
	branch := GetBranch()
	t.Error("===", branch)
}

func TestRun2(t *testing.T) {
	t.Log("testrun2")
	//t.Error("TestRun2")
}
