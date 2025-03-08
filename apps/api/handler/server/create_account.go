package server

import (
	"fmt"

	"github.com/byvinesse/vinance-backend/model"
	"github.com/byvinesse/vinance-backend/pkg/errors"
	"github.com/byvinesse/vinance-backend/pkg/response"
	"github.com/byvinesse/vinance-backend/pkg/validator"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateAccount(c echo.Context) error {
	ctx := c.Request().Context()
	userEmail := c.Get("user_email")
	userID := c.Get("user_id")

	var req model.CreateAccountRequest
	if err := c.Bind(&req); err != nil {
		return errors.ErrParseFailed(err)
	}

	if err := validator.ValidateStruct(ctx, req); err != nil {
		return errors.ErrParseFailed(err)
	}

	if req.Name == "" {
		return errors.ErrMissingField("name")
	} else if req.Currency == "" {
		return errors.ErrMissingField("currency")
	} else if req.Type == "" {
		return errors.ErrMissingField("type")
	} else if req.Color == "" {
		return errors.ErrMissingField("color")
	}

	res, err := h.app.AccountService.CreateAccount(ctx, userID.(string), &req)
	if err != nil {
		return errors.ErrInternalServerError(err, fmt.Sprintf("failed to #createAccount for request: %s of user %s", req.Name, userEmail))
	}

	return response.OkCreated(c, res)

}
