package routes

import (
	"github.com/labstack/gommon/log"
	"go-article/constant"
	"go-article/controller"
	"net/http"
)

type Routing struct {
	routes controller.CommonController
}

func (Routing Routing) GetRoutes() {

	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			Routing.routes.AddArticle(w, r)
		}
		if r.Method == "GET" {
			Routing.routes.GetArticle(w, r)
		}
	})

	log.Error(http.ListenAndServe(constant.Port, nil))
}
