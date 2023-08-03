package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
// var (
// 	ginDispatcher = gin.New()
// )

type ginRouter struct {
	ginDispatcher *gin.Engine
}

func (router *ginRouter) GET(uri string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.ginDispatcher.GET(uri, func(c *gin.Context) {
		handler(c.Writer, c.Request)
	})
}

func (router *ginRouter) POST(uri string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.ginDispatcher.POST(uri, func(c *gin.Context) {
		handler(c.Writer, c.Request)
	})
}

func (router *ginRouter) Serve(port string) error {
	return router.ginDispatcher.Run(fmt.Sprintf(":%s", port))
}

func NewGinRouter(ginDispatcher *gin.Engine) Router {
	return &ginRouter{
		ginDispatcher: ginDispatcher,
	}
}