package server

import (
	"fmt"

	"github.com/byvinesse/vinance-backend/pkg/errors"
	"github.com/byvinesse/vinance-backend/pkg/response"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetCompleteCategories(c echo.Context) error {
	ctx := c.Request().Context()
	userID := c.Get("user_id")
	email := c.Get("user_email")

	res, err := h.app.CategoryService.GetCompleteCategory(ctx, userID.(string))
	if err != nil {
		return errors.ErrInternalServerError(err, fmt.Sprintf("failed to #getCompleteCategories for user %s", email))
	}

	return response.Ok(c, res)
}
