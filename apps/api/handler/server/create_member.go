package server

import (
	"fmt"

	"github.com/byvinesse/vinance-backend/model"
	"github.com/byvinesse/vinance-backend/pkg/errors"
	"github.com/byvinesse/vinance-backend/pkg/response"
	"github.com/byvinesse/vinance-backend/pkg/validator"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateMember(c echo.Context) error {
	ctx := c.Request().Context()
	memberID := c.Get("member_id")
	memberEmail := c.Get("member_email")

	var req model.CreateMemberRequest
	if err := c.Bind(&req); err != nil {
		return errors.ErrParseFailed(err)
	}

	if err := validator.ValidateStruct(ctx, req); err != nil {
		return errors.ErrParseFailed(err)
	}
	if req.Username == "" {
		return errors.ErrMissingField("email")
	} else if req.PhoneNumber == "" {
		return errors.ErrMissingField("phone_number")
	} else if req.Gender == "" {
		return errors.ErrMissingField("gender")
	} else if req.DateOfBirth.IsZero() {
		return errors.ErrMissingField("date_of_birth")
	}

	req.Email = memberEmail.(string)

	_, err := h.app.MemberService.CreateMember(ctx, &req)
	if err != nil {
		return errors.ErrInternalServerError(err, fmt.Sprintf("failed to #createMember for request: %s", memberEmail))
	}

	ok, err := h.app.AuthService.CompleteMemberOnboarding(ctx, &model.CompleteMemberOnboardingRequest{
		ID: memberID.(string),
	})
	if err != nil {
		return err
	}

	return response.Ok(c, ok)
}
