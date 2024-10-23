package server

import (
	"fmt"

	"github.com/byvinesse/vinance-backend/model"
	"github.com/byvinesse/vinance-backend/pkg/errors"
	"github.com/byvinesse/vinance-backend/pkg/response"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(c echo.Context) error {
	ctx := c.Request().Context()

	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return errors.ErrParseFailed(err)
	}

	if req.Email == "" {
		return errors.ErrMissingField("email")
	} else if req.Password == "" {
		return errors.ErrMissingField("password")
	}

	result, err := h.app.AuthService.Register(ctx, &req)
	if err != nil {
		return errors.ErrInternalServerError(err, fmt.Sprintf("failed to #register new member for request: %s", req.Email))
	}

	return response.Ok(c, result)
}
