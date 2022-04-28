package server

import (
	"api-server/api/router"
	"api-server/config"
	"api-server/internal"
	"fmt"
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

func (s *Server) Start() error {
	svAddress := fmt.Sprintf("%s:%d", s.cfgs.ServerHost, s.cfgs.ServerPort)
	route := router.NewRouter(s.engine)
	err := route.RouterHandler()
	err = s.engine.Run(svAddress)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	err := s.inter.Stop()
	return err
}
