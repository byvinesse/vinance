package server

import (
	"github.com/labstack/echo/v4"
	"github.com/vincentkdeli/vinance-backend/pkg/response"
)

func (h *Handler) GetSample(c echo.Context) error { // TODO: delete later
	return response.Ok(c, true)
}
