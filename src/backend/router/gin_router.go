package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ginDispatcher = gin.Default()
)

type ginRouter struct {
	ginDispatcher *gin.Engine
}

func (*ginRouter) GET(uri string, handler func(w http.ResponseWriter, r *http.Request)) {
	ginDispatcher.GET(uri, func(c *gin.Context) {
		handler(c.Writer, c.Request)
	})
}

func (*ginRouter) POST(uri string, handler func(w http.ResponseWriter, r *http.Request)) {
	ginDispatcher.POST(uri, func(c *gin.Context) {
		handler(c.Writer, c.Request)
	})
}

func (*ginRouter) Serve(port string) error {
	return ginDispatcher.Run(fmt.Sprintf(":%s", port))
}

func NewGinRouter() Router {
	return &ginRouter{}
}
