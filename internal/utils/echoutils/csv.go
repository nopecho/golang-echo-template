package echoutils

import (
	"encoding/csv"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

type CsvData [][]string

func CsvResponse(c echo.Context, name string, data CsvData) error {
	disposition := fmt.Sprintf("attachment; filename=%s.csv", name)
	c.Response().Header().Set(echo.HeaderContentType, "text/csv")
	c.Response().Header().Set(echo.HeaderContentDisposition, disposition)

	writer := csv.NewWriter(c.Response().Writer)
	for _, row := range data {
		if err := writer.Write(row); err != nil {
			log.Warn().Msgf("error writing record to csv: %v, err: %v", row, err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		_ = c.String(http.StatusInternalServerError, fmt.Sprintf("error writing csv: %v", err))
	}
	return c.NoContent(http.StatusOK)
}

func NewCsvData(header []string) CsvData {
	return [][]string{header}
}

func AppendCsvData(data CsvData, row []string) CsvData {
	return append(data, row)
}
