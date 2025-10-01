package server

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/server/routes"
)

// Server 封装 Echo 实例和地址
type Server struct {
	Echo *echo.Echo
	Addr string
}

// NewServer 创建新的服务器实例
func NewServer(addr string) *Server {
	e := echo.New()
	routes.RegisterRoutes(e)

	return &Server{
		Echo: e,
		Addr: addr,
	}
}

// Start 启动服务器
func (s *Server) Start() error {
	fmt.Printf("Listening and serving HTTP on %s\n", s.Addr)
	return s.Echo.Start(s.Addr)
}
