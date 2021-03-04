###
 # @Date: 2020-11-09 13:42:04
 # @Author: fenggq
 # @LastEditors: fenggq
 # @LastEditTime: 2021-03-04 19:23:44
 # @FilePath: /godemo/deploy/deploy.sh
### 
#! /bin/bash
#

declare -A gitNameMap
gitNameMap["andyfenggq"]="冯国庆"
gitNameMap["Aleutian Xie"]="谢辉生"
gitNameMap["luo"]="骆玉霞"
gitNameMap["audu"]="杜于庆"

declare -A gitEmailMap
gitEmailMap["362739259@qq.com"]="冯国庆"
gitEmailMap["aleutian.xie@cicisoft.cn"]="谢辉生"
gitEmailMap["luo_yu_xia@163.com"]="骆玉霞"
gitEmailMap["audu@qq.com"]="杜于庆"
declare -A dingding
dingding["冯国庆"]=17316225231
dingding["骆玉霞"]=13552079799
dingding["谢辉生"]=15901435695
dingding["姜亦春"]=13581894261
dingding["杜于庆"]=18211025188


function strinclude() {
  s1=$1
  s2=$2
  echo $s1 == $s2
  result=$(echo $s1 | grep "${s2}")
  if [[ "$result" != "" ]]
  then
      echo "$s1 include $s2"
  else
      echo "$1 not include $s2"
  fi
}
GIT_COMMIT_NAME=$(git show -s --format='%cn')
GIT_COMMIT_EMAIL=$(git show -s --format='%ce')
GIT_COMMIT_TIME=$(git show -s --format='%cd')
GIT_COMMIT_MESSAGE=$(git show -s --format='%s')
GIT_BRANCH=$(git name-rev --name-only HEAD)
BUILD_VERSION=$(git log -1 --oneline)
BUILD_TIME=$(date "+%FT%T%z")
APP_VERSION=$(git describe --abbrev=0)
GO_VERSION=$(go version)

echo GIT_COMMIT_NAME $GIT_COMMIT_NAME
echo GIT_COMMIT_EMAIL $GIT_COMMIT_EMAIL
echo GIT_COMMIT_TIME $GIT_COMMIT_TIME
echo GIT_COMMIT_MESSAGE $GIT_COMMIT_MESSAGE
echo GIT_BRANCH $GIT_BRANCH
echo BUILD_VERSION $BUILD_VERSION
echo BUILD_TIME $BUILD_TIME
echo APP_VERSION $APP_VERSION

if [[ "$GIT_COMMIT_NAME" != "" ]]
then
    GIT_COMMIT_USER_NAME=${gitNameMap[$GIT_COMMIT_NAME]}
    if [[ "$GIT_COMMIT_USER_NAME" == "" ]]
    then
        GIT_COMMIT_USER_NAME=$GIT_COMMIT_NAME
    fi
else
    if [[ "$GIT_COMMIT_EMAIL" != "" ]]
    then
        GIT_COMMIT_USER_NAME=${gitEmailMap[$GIT_COMMIT_EMAIL]}
        if [[ "$Name" == "" ]]
        then
            GIT_COMMIT_USER_NAME=$GIT_COMMIT_EMAIL
        fi
    else
       echo GIT_COMMIT_USER_NAME="name and email"
    fi
fi
echo $GIT_COMMIT_USER_NAME

if [[ "$GIT_COMMIT_USER_NAME" != "" ]]
then
    DINGDING_PHONE=${dingding[$GIT_COMMIT_USER_NAME]}
else
    GIT_COMMIT_USER_NAME="name and email is null"
fi