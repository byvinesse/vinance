package server

import (
	"fmt"
	"time"

	"github.com/byvinesse/vinance-backend/model"
	"github.com/byvinesse/vinance-backend/pkg/errors"
	"github.com/byvinesse/vinance-backend/pkg/response"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateRecord(c echo.Context) error {
	ctx := c.Request().Context()
	userID := c.Get("user_id")
	userEmail := c.Get("user_email")

	var req model.CreateRecordRequest
	if err := c.Bind(&req); err != nil {
		return errors.ErrParseFailed(err)
	}

	if req.AccountID == "" {
		return errors.ErrMissingField("account_id")
	} else if req.SubCategoryID == "" {
		return errors.ErrMissingField("subcategory_id")
	} else if req.Amount <= 0 {
		return errors.ErrInvalidValue("amount")
	} else if req.Currency == "" {
		return errors.ErrMissingField("currency")
	} else if req.BaseAmount <= 0 {
		return errors.ErrInvalidValue("base_amount")
	} else if req.Type == "" {
		return errors.ErrMissingField("type")
	} else if req.PaymentType == "" {
		return errors.ErrMissingField("payment_type")
	} else if req.PaymentStatus == "" {
		return errors.ErrMissingField("payment_status")
	}

	if req.RecordedAt == nil {
		now := time.Now()
		req.RecordedAt = &now
	}

	res, err := h.app.RecordService.CreateRecord(ctx, userID.(string), &req)
	if err != nil {
		return errors.ErrInternalServerError(err, fmt.Sprintf("failed to #createRecord for user %s", userEmail))
	}

	return response.OkCreated(c, res)
}
