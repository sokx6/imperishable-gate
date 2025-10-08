package server

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"imperishable-gate/internal/server/database"
	"imperishable-gate/internal/server/routes"
	"imperishable-gate/internal/server/utils/logger"
	"imperishable-gate/internal/types/data"
)

// Server 封装 Echo 实例和地址
type Server struct {
	Echo *echo.Echo
	Addr string
}

// NewServer 创建新的服务器实例
func NewServer(addr, dsn string) *Server {
	if addr == "" {
		addr = "localhost:4514"
	}
	if dsn == "" {
		dsn = "host=localhost user=postgres dbname=gate_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	}
	// 创建新的echo实例
	e := echo.New()

	// 初始化日志器
	logger.InitLogger(e)

	// 添加请求ID
	e.Use(middleware.RequestID())

	// 自定义日志中间件：带颜色和用户信息
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			// 先执行下一个中间件/处理器（包括JWT中间件）
			err := next(c)

			// 在这里，userInfo已经被JWT中间件设置了
			// 获取请求信息
			req := c.Request()
			res := c.Response()
			latency := time.Since(start)

			// 获取用户信息（现在可以正确获取了）
			userInfo := "Guest"
			userColor := "\033[90m" // 灰色 - 游客
			userInfoValue := c.Get("userInfo")
			if userInfoValue != nil {
				if user, ok := userInfoValue.(data.UserInfo); ok {
					userInfo = fmt.Sprintf("User: %s (ID:%d)", user.Username, user.UserID)
					userColor = "\033[34m" // 蓝色 - 已认证用户
				}
			}

			// 状态码颜色
			statusColor := getStatusColor(res.Status)

			// 构建日志
			log := fmt.Sprintf(
				"%s%03d\033[0m | \033[33m%-6s\033[0m \033[32m%-40s\033[0m | \033[35m%10v\033[0m | %-15s | %s%s\033[0m",
				statusColor,
				res.Status,
				req.Method,
				req.RequestURI,
				latency.Round(time.Microsecond),
				c.RealIP(),
				userColor,
				userInfo,
			)

			fmt.Println(log)
			return err
		}
	})

	e.Use(middleware.Recover())

	// 初始化数据库
	if err := database.InitDB(dsn); err != nil {
		e.Logger.Fatal("Failed to connect to database: ", err)
	}

	// 注册路由
	routes.RegisterRoutes(e)

	// 返回Server实例
	return &Server{
		Echo: e,
		Addr: addr,
	}
}

// Start 启动服务器
func (s *Server) Start() error {
	logger.Info("Listening and serving HTTP on %s", s.Addr)
	return s.Echo.Start(s.Addr)
}

// getStatusColor 根据状态码返回颜色代码
func getStatusColor(status int) string {
	switch {
	case status >= 200 && status < 300:
		return "\033[32m" // 绿色 - 成功
	case status >= 300 && status < 400:
		return "\033[36m" // 青色 - 重定向
	case status >= 400 && status < 500:
		return "\033[33m" // 黄色 - 客户端错误
	case status >= 500:
		return "\033[31m" // 红色 - 服务器错误
	default:
		return "\033[0m" // 默认
	}
}
