package echoserver

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nopecho/golang-template/pkg/echoserver/e"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewEcho() (e *echo.Echo) {
	e = echo.New()
	defaultMiddleware(e)
	e.GET("/health", healthHandler)
	return e
}

// Run starts the echo server with graceful shutdown
func Run(e *echo.Echo, port string) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer stop()
	go startServer(e, port)
	<-ctx.Done()
	stopServer(e)
}

func startServer(e *echo.Echo, port string) {
	err := e.Start(fmt.Sprintf(":%s", port))
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		e.Logger.Fatal(err)
	}
}

func stopServer(e *echo.Echo) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := e.Shutdown(ctx)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		e.Logger.Fatal(err)
	}
}

func defaultMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, e.Map{
		"status": "ok",
	})
}
