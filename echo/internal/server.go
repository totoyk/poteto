package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/totoyk/echo/internal/infra/db"
	"github.com/totoyk/echo/internal/interfaces"
	"github.com/totoyk/echo/internal/interfaces/oas"
)

func Run() error {
	e := echo.New()
	g := e.Group("/v1")

	g.Use(middleware.Logger())
	g.Use(middleware.Recover())

	dbconn, err := db.NewConn()
	if err != nil {
		return err
	}
	handlers := interfaces.NewHandlers(dbconn)
	oas.RegisterHandlers(g, handlers)

	e.Logger.Fatal(e.Start(":1323"))

	return nil
}
