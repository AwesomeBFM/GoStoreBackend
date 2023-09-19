package router

import "github.com/gin-gonic/gin"

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Start() error {
	router := gin.Default()

	return router.Run(":8080")
}
