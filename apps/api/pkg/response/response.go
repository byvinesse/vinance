package response

import (
	"net/http"

	"github.com/byvinesse/vinance-backend/entity"
	"github.com/labstack/echo/v4"
)

func Ok[T any](c echo.Context, data T) error {
	return c.JSON(http.StatusOK, entity.OkResponse[T]{
		Code:   200,
		Status: "OK",
		Data:   data,
	})
}
