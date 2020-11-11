/*
 * @Date: 2020-11-10 20:32:43
 * @Author: fenggq
 * @LastEditors: fenggq
 * @LastEditTime: 2020-11-11 19:15:12
 * @FilePath: /godemo/gitlog.go
 */
package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

//GitLog ...
type GitLog struct {
	Commit  string
	Author  string
	Date    string
	Message string
}

var glogs []GitLog

//LoadLatestGitLogs ...
func LoadLatestGitLogs() string {
	gitcmd := "git log -1"
	gitlog, err := CMD(gitcmd)
	if err != nil {
		log.Println(err)
		return ""
	}
	return gitlog
}

//LoadLatestCommitUser ...
func LoadLatestCommitUser() string {
	gitcmd := "git show -s --format='%cn'"
	gitlog, err := CMD(gitcmd)
	if err != nil {
		log.Println(err)
		return ""
	}
	return strings.TrimSpace(gitlog)
}

//LoadCommitMessage ...
func LoadCommitMessage(commit string) string {
	gitcmd := "git log --pretty=oneline "
	if commit != "" {
		gitcmd = fmt.Sprintf("%s %s...", gitcmd, commit)
	} else {
		nowDate := time.Now()
		nowDate = nowDate.AddDate(0, 0, -2)
		//log.Println(nowDate)
		gitcmd = fmt.Sprintf("%s --after=%s", gitcmd, nowDate.Format("2006-01-02"))
	}
	log.Println(gitcmd)
	gitlog, err := CMD(gitcmd)
	if err != nil {
		log.Println(err)
		return ""
	}
	//LoadCommitMessageParse(gitlog)
	return gitlog
}

//LoadCommitMessageParse ...
func LoadCommitMessageParse(Message string) {
	array := strings.Split(Message, "\n")
	for k, v := range array {
		msg := strings.SplitN(v, " ", 2)
		for km, vm := range msg {
			log.Println(k, km, vm)
		}

	}
	/*array := strings.SplitN(Message, "\n", 1)
	for k, v := range array {
		log.Println(k, v)
	}*/

}

//LoadGitLogs ...
func LoadGitLogs(commit string) string {
	gitcmd := "git log --no-merges --date=format:'%Y-%m-%d %H:%M:%S'"
	if commit != "" {
		gitcmd = fmt.Sprintf("%s %s...", gitcmd, commit)
	} else {
		nowDate := time.Now()
		nowDate = nowDate.AddDate(0, 0, -7)
		log.Println(nowDate)
		gitcmd = fmt.Sprintf("%s --after=%s", gitcmd, nowDate.Format("2006-01-02"))
	}
	log.Println(gitcmd)
	gitlog, err := CMD(gitcmd)
	if err != nil {
		log.Println(err)
		return ""
	}
	return gitlog
}

//LoadGitLog ...
func LoadGitLog(commit string) {
	gitcmd := "git log --no-merges --date=format:'%Y-%m-%d %H:%M:%S'"
	if commit != "" {
		gitcmd = fmt.Sprintf("%s %s...", gitcmd, commit)
	} else {
		nowDate := time.Now()
		nowDate = nowDate.AddDate(0, 0, -7)
		log.Println(nowDate)
		gitcmd = fmt.Sprintf("%s --after=%s", gitcmd, nowDate.Format("2006-01-02"))
	}
	log.Println(gitcmd)
	gitlog, err := CMD(gitcmd)
	if err != nil {
		log.Println(err)
		return
	}
	kv := strings.Split(gitlog, "\n")
	for k, v := range kv {
		log.Println(k, v)
	}
	kvLen := len(kv)

	slen := kvLen / 6
	log.Println(slen)
	subNum := 0
	for i := 0; i < slen; i++ {
		glog := GitLog{}
		num := i*6 + subNum
		commit := strings.Split(kv[0+num], " ")
		author := strings.Split(kv[1+num], ":")
		date := strings.Split(kv[2+num], ": ")

		log.Println(num, commit, author, date)
		glog.Commit = strings.TrimSpace(commit[1])
		glog.Author = strings.TrimSpace(author[1])
		glog.Date = strings.TrimSpace(date[1])
		glog.Message = strings.TrimSpace(kv[4+num])
		log.Println(glog.Message)
		for mi := 5 + num; mi < kvLen-2; mi++ {
			log.Println(mi, len(kv[mi]), kv[mi])
			if len(kv[mi]) > 0 {
				subNum++
				log.Println(subNum)
				glog.Message += strings.TrimSpace(kv[4+num])
			}
			if len(kv[mi+1]) == 0 {
				log.Println("===", subNum)
				break
			}
		}
		glogs = append(glogs, glog)
	}
	log.Println(glogs)
}

//SaveCommit ...
func SaveCommit() {

}
