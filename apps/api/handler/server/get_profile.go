package server

import (
	"fmt"

	"github.com/byvinesse/vinance-backend/pkg/errors"
	"github.com/byvinesse/vinance-backend/pkg/response"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetProfile(c echo.Context) error {
	ctx := c.Request().Context()
	email := c.Get("user_email")

	res, err := h.app.UserService.GetProfile(ctx, email.(string))
	if err != nil {
		return errors.ErrInternalServerError(err, fmt.Sprintf("failed to #getProfile for user %s", email))
	}

	return response.Ok(c, res)
}
