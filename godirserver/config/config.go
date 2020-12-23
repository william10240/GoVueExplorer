package config

import (
	"os"
	"path"
)

// 定义当前目录
var APPPATH string
var BASEPATH string

func init() {
	p, _ := os.Getwd()
	APPPATH = p
	if len(os.Args) > 1 {
		BASEPATH = path.Join(APPPATH,os.Args[1])
	}else{
		BASEPATH = path.Join(APPPATH,"images")
	}
}



