package zcache

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/armnerd/zcache/internal/route"
)

type Server struct {
	zinx ziface.IServer
}

func NewServer() *Server {
	s := &Server{}
	z := znet.NewServer()
	z.AddRouter(0, &route.Router{})
	s.zinx = z
	return s
}

func (s *Server) Run() {
	s.zinx.Serve()
}
