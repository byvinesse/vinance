package errors

import (
	"fmt"
	"net/http"

	"github.com/byvinesse/vinance-backend/entity"
	"github.com/byvinesse/vinance-backend/pkg/log"
	"github.com/labstack/echo/v4"
)

func Middleware(err error, c echo.Context) {
	// 400 Bad Request Error
	var badRequestError *BadRequestError
	if As(err, &badRequestError) {
		log.Error().Err(err).Caller().Msg(fmt.Sprintf("app_id=%s | member_email=%s | error=%v", c.Get("app_id"), c.Get("member_email"), badRequestError.InternalError))
		_ = c.JSON(http.StatusBadRequest, entity.ErrResponse{
			Code:    badRequestError.Code,
			Status:  badRequestError.Status,
			Message: badRequestError.Message,
		})
		return
	}

	// 400 Client Error
	var validationError *ValidationError
	if As(err, &validationError) {
		log.Error().Err(err).Caller().Msg(fmt.Sprintf("app_id=%s | member_email=%s | error=%v", c.Get("app_id"), c.Get("member_email"), validationError.InternalError))
		_ = c.JSON(http.StatusBadRequest, entity.ErrResponse{
			Code:    validationError.Code,
			Status:  validationError.Status,
			Message: validationError.Message,
		})
		return
	}

	// 401 Authorization Error
	var unauthorizedError *UnauthorizedError
	if As(err, &unauthorizedError) {
		log.Error().Err(err).Caller().Msg(fmt.Sprintf("app_id=%s | member_email=%s | error=%v", c.Get("app_id"), c.Get("member_email"), unauthorizedError.InternalError))
		_ = c.JSON(http.StatusUnauthorized, entity.ErrResponse{
			Code:    unauthorizedError.Code,
			Status:  unauthorizedError.Status,
			Message: unauthorizedError.Message,
		})
		return
	}

	// 403 Forbidden Error
	var forbiddenError *ForbiddenError
	if As(err, &forbiddenError) {
		log.Error().Err(err).Caller().Msg(fmt.Sprintf("app_id=%s | member_email=%s | error=%v", c.Get("app_id"), c.Get("member_email"), forbiddenError.InternalError))
		_ = c.JSON(http.StatusForbidden, entity.ErrResponse{
			Code:    forbiddenError.Code,
			Status:  forbiddenError.Status,
			Message: forbiddenError.Message,
		})
		return
	}

	// 404 Not Found Error
	var notFoundError *NotFoundError
	if As(err, &notFoundError) {
		log.Error().Err(err).Caller().Msg(fmt.Sprintf("app_id=%s | member_email=%s | error=%v", c.Get("app_id"), c.Get("member_email"), notFoundError.InternalError))
		_ = c.JSON(http.StatusNotFound, entity.ErrResponse{
			Code:    notFoundError.Code,
			Status:  notFoundError.Status,
			Message: notFoundError.Message,
		})
		return
	}

	// 409 Duplicate Error
	var duplicateError *DuplicateError
	if As(err, &duplicateError) {
		log.Error().Err(err).Caller().Msg(fmt.Sprintf("app_id=%s | member_email=%s | error=%v", c.Get("app_id"), c.Get("member_email"), duplicateError.InternalError))
		_ = c.JSON(http.StatusConflict, entity.ErrResponse{
			Code:    duplicateError.Code,
			Status:  duplicateError.Status,
			Message: duplicateError.Message,
		})
		return
	}

	// 500 Internal Server Error
	var serverError *ServerError
	if As(err, &serverError) {
		log.Error().Err(err).Caller().Msg(fmt.Sprintf("app_id=%s | member_email=%s | error=%v", c.Get("app_id"), c.Get("member_email"), serverError.InternalError))
		_ = c.JSON(http.StatusInternalServerError, entity.ErrResponse{
			Code:    serverError.Code,
			Status:  serverError.Status,
			Message: "Sorry, something went wrong.",
		})
		return
	}

	// Unknown Error
	log.Error().Err(err).Caller().Msg(fmt.Sprintf("app_id=%s | member_email=%s | error=%v", c.Get("app_id"), c.Get("member_email"), err))
	_ = c.JSON(http.StatusBadGateway, entity.ErrResponse{
		Code:    502,
		Status:  "INTERNAL_SERVER_ERROR",
		Message: "Sorry, something went wrong",
	})
}
