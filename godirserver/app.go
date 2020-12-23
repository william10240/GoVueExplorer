package main

import (
	"./router"
	"log"
	"net/http"
)

func main() {

	rut := router.GetRouter()
	err := http.ListenAndServe(":8000", rut)
	if err != nil{
		log.Fatal(err)
	}else{
		log.Println("Server started at http://127.0.0.1:8000")
	}
	//scanDir(APP_PATH)
}
