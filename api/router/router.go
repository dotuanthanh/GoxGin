package router

import (
	"api-server/api/delivery/admin"
	"github.com/gin-gonic/gin"
)

type router struct {
	engine *gin.Engine
}

func NewRouter(engine *gin.Engine) *router {
	return &router{
		engine: engine,
	}
}
func (r *router) RouterHandler() error {
	r.engine.POST("/sign-up", admin.SignUp)
	r.engine.POST("/log-in", admin.Login)

	return nil
}
