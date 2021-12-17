package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddlewareInit(e *echo.Echo) {
	logger := middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}",` +
    			`"method":"${method}","host":"${host}","uri":"${uri}",` +
    			`"status":${status},"error":"${error}","latency":${latency_human}}` + "\n",
	}

	e.Use(middleware.LoggerWithConfig(logger))
}
