package echo

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewEcho() (e *echo.Echo) {
	e = echo.New()
	defaultMiddleware(e)
	defaultRoute(e)
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

func defaultRoute(e *echo.Echo) {
	e.GET("/", noContentHandler)
	e.GET("/favicon.ico", noContentHandler)
	e.GET("/health", healthHandler)
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, MAP{
		"status": "ok",
	})
}

func noContentHandler(c echo.Context) error {
	return c.JSON(http.StatusNoContent, nil)
}
