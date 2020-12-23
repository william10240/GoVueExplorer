package config

import (
	"log"
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
		BASEPATH = path.Join(APPPATH, os.Args[1])
	} else {
		BASEPATH = path.Join(APPPATH, "images")
	}
	s, err := os.Stat(BASEPATH)
	if err != nil {
		err = os.Mkdir(BASEPATH, os.ModeDir)
		if err != nil {
			log.Fatal("Mkdir BASEPATH fail.")
		}
	} else if s != nil && s.IsDir() == false {
		err = os.Remove(BASEPATH)
		if err != nil {
			log.Fatal("Remove BASEPATH fail")
		}
		
		err = os.Mkdir(BASEPATH, os.ModeDir)
		if err != nil {
			log.Fatal("Mkdir BASEPATH fail")
		}
		
	}
}
