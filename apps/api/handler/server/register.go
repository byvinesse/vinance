package server

import (
	"github.com/labstack/echo/v4"
	"github.com/vincentkdeli/vinance-backend/model"
	"github.com/vincentkdeli/vinance-backend/pkg/errors"
	"github.com/vincentkdeli/vinance-backend/pkg/response"
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
	} else if req.PhoneNumber == "" {
		return errors.ErrMissingField("phone_number")
	}

	result, err := h.app.AuthService.Register(ctx, &req)
	if err != nil {
		return errors.ErrInternalServerError(err, "failed to register new member")
	}

	return response.Ok(c, result)
}
