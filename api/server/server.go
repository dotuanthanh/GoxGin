package server

import (
	"api-server/config"
	"api-server/internal"
	"github.com/gin-gonic/gin"
)

type Server struct {
	inter  *internal.Internal
	cfgs   *config.Server
	engine *gin.Engine
}

func NewServer(inter *internal.Internal, cfgs *config.Server, engine *gin.Engine) *Server {
	return &Server{
		inter:  inter,
		cfgs:   cfgs,
		engine: engine,
	}
}

func (s *Server) Start() {

}

func (s *Server) Stop() {

}
