/*
 * @Date: 2020-11-08 19:04:14
 * @Author: fenggq
 * @LastEditors: fenggq
 * @LastEditTime: 2020-11-08 19:50:03
 * @FilePath: /godemo/main.go
 */
package godemo

import "log"

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

//
func main() {
	log.SetFlags(log.Lshortfile)
	log.Printf("App Name:\t%s\n", AppName)
	log.Printf("App Version:\t%s\n", AppVersion)
	log.Printf("Build version:\t%s\n", BuildVersion)
	log.Printf("Build time:\t%s\n", BuildTime)
	log.Printf("Git revision:\t%s\n", GitRevision)
	//log.Printf("Git branch:\t%s\n", GitBranch)
	log.Printf("Golang Version: %s\n", GoVersion)
	log.Println("=====")
}
