package utils

import (
	"../config"
	"fmt"
	"io/ioutil"
	"path"
)

func ScanDir(dirPath string) ([]string, []string) {
	realPath := path.Join(config.BASEPATH, dirPath)
	dir, err := ioutil.ReadDir(realPath)
	if err != nil {
		fmt.Println(err)
	}
	dirs := make([]string, 0, 10)
	files := make([]string, 0, 10)
	for _, fi := range dir {
		if fi.IsDir() {
			dirs = append(dirs, path.Join(dirPath, fi.Name()))
		} else {
			files = append(files, path.Join(dirPath, fi.Name()))
		}
	}
	return dirs, files
}
