package model


type JsonResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data interface{}    `json:"data"`
}
type Res struct {
	Dirs  []string `json:"dirs"`
	Files []string `json:"files"`
}
