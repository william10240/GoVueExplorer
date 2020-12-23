package router

import (
	"../controller"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var router = httprouter.New()

func init() {
fmt.Println("route init")
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})
}

func GetRouter() *httprouter.Router {

	router.GET("/", controller.Index)
	router.POST("/u", controller.Upload)
	router.GET("/d", controller.Delete)
	return router
}
