package route

import (
	"github.com/eurie-inc/echo-sample/api"
	"github.com/eurie-inc/echo-sample/db"
	"github.com/eurie-inc/echo-sample/handler"
	myMw "github.com/eurie-inc/echo-sample/middleware"
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Debug()

	// Set Bundle MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// Set Custom MiddleWare
	e.Use(myMw.TransactionHandler(db.Init()))

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.POST("/members", api.PostMember())
		v1.GET("/members", api.GetMembers())
		v1.GET("/members/:id", api.GetMember())
	}
	return e
}
