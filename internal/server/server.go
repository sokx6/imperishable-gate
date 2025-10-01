package server

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/routes"
)

const defaultDSN = "host=localhost user=postgres dbname=gate_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"

// Server 封装 Echo 实例和地址
type Server struct {
	Echo *echo.Echo
	Addr string
}

// NewServer 创建新的服务器实例
func NewServer(addr, dsn string) *Server {
	if dsn == "" {
		dsn = defaultDSN
	}

	e := echo.New()

	if err := database.InitDB(dsn); err != nil {
		e.Logger.Fatal("Failed to connect to database: ", err)
	}

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
