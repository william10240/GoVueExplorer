package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type JsonResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Res    `json:"data"`
}
type Res struct {
	Dirs  []string `json:"dirs"`
	Files []string `json:"files"`
}

// 定义当前目录
var APPPATH string
func init(){
	p:=os.Args[1]
	if p == "" {
		p,_ = os.Getwd()
	}
	APPPATH = p
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	p := r.FormValue("p")
	if p == "" {
		p = "/"
	}
	p = strings.Replace(p,"../","/",-1)
	s, _ := os.Stat(path.Join(APPPATH, p))    //os.Stat获取文件信息
	if(s.IsDir()){
		dirs, files := scanDir(p)
		resData := JsonResult{Code: 200, Msg: "", Data: Res{Dirs: dirs, Files: files}}

		msg, _ := json.Marshal(resData)

		//fmt.Fprintf(w, "hello world")
		w.Header().Set("content-type", "text/json")
		w.Write(msg)
	}else{
		//w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", path.Base(p)))
		fData,_ := ioutil.ReadFile(path.Join(APPPATH, p))
		w.Write(fData)
	}

}
func uploadHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm() //解析表单
	imgFile, _, err := r.FormFile("f")//获取文件内容
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()
	imgName := ""
	files := r.MultipartForm.File //获取表单中的信息
	for k, v := range files {
		for _, vv := range v {
			fmt.Println(k + ":" + vv.Filename)//获取文件名
			if strings.Index(vv.Filename, ".png") > 0 {
				imgName = vv.Filename
			}
		}
	}
	p := r.FormValue("p")
	if p == "" {
		p = "/"
	}
	p = strings.Replace(p,"../","/",-1)
	p = path.Join(APPPATH, p)
	imgPath := path.Join(p,imgName)


	saveFile, _ := os.Create(imgPath)
	defer saveFile.Close()
	s,err := io.Copy(saveFile, imgFile) //保存
	fmt.Println(s,err)

	resData := JsonResult{Code: 200, Msg: ""}
	msg, _ := json.Marshal(resData)
	w.Write(msg)
}
func deleteHandler(w http.ResponseWriter, r *http.Request){
	p := r.FormValue("p")
	if p == "" {p = "/"}
	p = strings.Replace(p,"../","/",-1)
	p = path.Join(APPPATH, p)


	err := os.Remove(p)
	fmt.Println(err)

	resData := JsonResult{Code: 200, Msg: ""}
	msg, _ := json.Marshal(resData)
	w.Write(msg)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/u", uploadHandler)
	http.HandleFunc("/d", deleteHandler)
	http.ListenAndServe(":8000", nil)
	//scanDir(APP_PATH)
}

func scanDir(dirPath string) ([]string, []string) {
	realPath := path.Join(APPPATH, dirPath)
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
