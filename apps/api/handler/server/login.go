package server

import (
	"fmt"

	"github.com/byvinesse/vinance-backend/model"
	"github.com/byvinesse/vinance-backend/pkg/errors"
	"github.com/byvinesse/vinance-backend/pkg/response"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return errors.ErrParseFailed(err)
	}

	if req.Email == "" {
		return errors.ErrMissingField("email")
	} else if req.Password == "" {
		return errors.ErrMissingField("password")
	}

	result, err := h.app.UserService.Login(ctx, &req)
	if err != nil {
		return errors.ErrInternalServerError(err, fmt.Sprintf("failed to #login user for request: %s", req.Email))
	}

	return response.Ok(c, result)
}
