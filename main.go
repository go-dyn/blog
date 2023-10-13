package main

import (
	"net/http"

	"github.com/go-dyn/blog/app/controller"
	"github.com/syumai/workers"
)

func main() {
	//http.Handle("/categories", controller.NewCategoryHandler())
	//http.Handle("/tags", controller.NewTagHandler())
	http.Handle("/users", controller.NewUserHandler())
	//http.Handle("/articles", controller.NewPostHandler())
	//http.Handle("/comments", controller.NewCommentHandler())
	workers.Serve(nil) // use http.DefaultServeMux
}
