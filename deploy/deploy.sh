###
 # @Date: 2020-11-09 13:42:04
 # @Author: fenggq
 # @LastEditors: fenggq
 # @LastEditTime: 2020-11-09 13:43:01
 # @FilePath: /godemo/deploy/deploy.sh
### 
#! /bin/bash

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