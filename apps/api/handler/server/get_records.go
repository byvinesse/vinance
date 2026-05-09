package server

import (
	"fmt"
	"strconv"

	"github.com/byvinesse/vinance-backend/pkg/errors"
	"github.com/byvinesse/vinance-backend/pkg/response"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetRecords(c echo.Context) error {
	ctx := c.Request().Context()
	userID := c.Get("user_id")
	userEmail := c.Get("user_email")

	cursorStr := c.QueryParam("cursor")

	limit := 0
	if raw := c.QueryParam("limit"); raw != "" {
		parsed, err := strconv.Atoi(raw)
		if err != nil || parsed <= 0 {
			return errors.ErrInvalidValue("limit")
		}
		limit = parsed
	}

	res, err := h.app.RecordService.GetRecords(ctx, userID.(string), limit, cursorStr)
	if err != nil {
		return errors.ErrInternalServerError(err, fmt.Sprintf("failed to #getRecords for user %s", userEmail))
	}

	return response.Ok(c, res)
}
