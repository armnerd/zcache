package zcache

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/armnerd/zcache/internal/expire"
	"github.com/armnerd/zcache/internal/land"
	"github.com/armnerd/zcache/internal/route"
)

type Server struct {
	zinx ziface.IServer
	opts *Options
}

func NewServer(opts ...OptionFunc) *Server {
	options := loadOptions(opts...)
	z := znet.NewServer()
	z.AddRouter(0, &route.Router{})
	s := &Server{
		zinx: z,
		opts: options,
	}
	return s
}

func (s *Server) Run() {
	go expire.Clean(s.opts.CleanSeq) // 清理过期数据
	go land.Land(s.opts.LandSeq)     // 定时持久化
	s.zinx.Serve()
}
