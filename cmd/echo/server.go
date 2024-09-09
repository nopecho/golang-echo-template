package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nopecho/golang-template/internal/app/config"
	"github.com/nopecho/golang-template/pkg/echoserver"
	"github.com/nopecho/golang-template/pkg/synk"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	e := echoserver.NewEcho()

	e.GET("goroutine", func(c echo.Context) error {
		format := echoserver.DefaultQuery(c, "format", "json")
		sizeParam := echoserver.DefaultQuery(c, "size", "1000")
		size, _ := strconv.Atoi(sizeParam)

		cm := synk.OpenBufferChannel[string](size)
		minN := 0
		maxN := 3
		go func() {
			for i := range cm.Size {
				go func() {
					time.Sleep(time.Duration(rand.Intn(maxN-minN+1)+minN) * time.Second)
					cm.Send(fmt.Sprintf("test %d sequence", i))
				}()
			}
		}()
		result := cm.WaitReceive()

		switch format {
		case "json":
			return c.JSON(http.StatusOK, result)
		case "csv":
			data := echoserver.NewCsvData([]string{"idx", "result"})
			for i, v := range result {
				data = echoserver.AppendCsvData(data, []string{strconv.Itoa(i), v})
			}
			return echoserver.CsvResponse(c, "goroutine", data)
		default:
			return c.JSON(http.StatusBadRequest, echoserver.Map{
				"error": "invalid format",
			})
		}
	})

	echoserver.Run(e, config.Env.Port)
}
