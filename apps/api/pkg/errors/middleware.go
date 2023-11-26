package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/vincentkdeli/vinance-backend/entity"
)

func Middleware(err error, c echo.Context) {
	// 400 Bad Request Error
	var badRequestError *BadRequestError
	if As(err, &badRequestError) {
		log.Printf("error=%v app_id=%s member_email=%s\n", badRequestError.InternalError, c.Get("app_id"), c.Get("member_email"))
		_ = c.JSON(http.StatusBadRequest, entity.ErrResponse{
			Status:    "BAD_REQUEST",
			ErrorCode: badRequestError.ErrorCode,
			Message:   badRequestError.Message,
		})
		return
	}

	// 400 Client Error
	var validationError *ValidationError
	if As(err, &validationError) {
		log.Printf("error=%v app_id=%s member_email=%s\n", validationError.InternalError, c.Get("app_id"), c.Get("member_email"))
		_ = c.JSON(http.StatusBadRequest, entity.ErrResponse{
			Status:    "VALIDATION_ERROR",
			ErrorCode: validationError.ErrorCode,
			Message:   validationError.Message,
		})
		return
	}

	// 401 Authorization Error
	var unauthorizedError *UnauthorizedError
	if As(err, &unauthorizedError) {
		log.Printf("error=%v app_id=%s member_email=%s\n", unauthorizedError.InternalError, c.Get("app_id"), c.Get("member_email"))
		_ = c.JSON(http.StatusUnauthorized, entity.ErrResponse{
			Status:    "UNAUTHORIZED",
			ErrorCode: unauthorizedError.ErrorCode,
			Message:   unauthorizedError.Message,
		})
		return
	}

	// 403 Forbidden Error
	var forbiddenError *ForbiddenError
	if As(err, &forbiddenError) {
		log.Printf("error=%v app_id=%s member_email=%s\n", forbiddenError.InternalError, c.Get("app_id"), c.Get("member_email"))
		_ = c.JSON(http.StatusForbidden, entity.ErrResponse{
			Status:    "FORBIDDEN",
			ErrorCode: forbiddenError.ErrorCode,
			Message:   forbiddenError.Message,
		})
		return
	}

	// 404 Not Found Error
	var notFoundError *NotFoundError
	if As(err, &notFoundError) {
		log.Printf("error=%v app_id=%s member_email=%s\n", notFoundError.InternalError, c.Get("app_id"), c.Get("member_email"))
		_ = c.JSON(http.StatusNotFound, entity.ErrResponse{
			Status:    "NOT_FOUND",
			ErrorCode: notFoundError.ErrorCode,
			Message:   notFoundError.Message,
		})
		return
	}

	// 409 Duplicate Error
	var duplicateError *DuplicateError
	if As(err, &duplicateError) {
		log.Printf("error=%v app_id=%s member_email=%s\n", notFoundError.InternalError, c.Get("app_id"), c.Get("member_email"))
		_ = c.JSON(http.StatusConflict, entity.ErrResponse{
			Status:    "DATA_ALREADY_EXISTS",
			ErrorCode: duplicateError.ErrorCode,
			Message:   duplicateError.Message,
		})
		return
	}

	// 500 Internal Server Error
	var serverError *ServerError
	if As(err, &serverError) {
		log.Printf("error=%v app_id=%s member_email=%s\n", serverError.InternalError, c.Get("app_id"), c.Get("member_email"))
		_ = c.JSON(http.StatusInternalServerError, entity.ErrResponse{
			Status:    "INTERNAL_SERVER_ERROR",
			ErrorCode: "INTERNAL_SERVER_ERROR",
			Message:   "Sorry, something went wrong",
		})
		return
	}

	// Unknown Error
	log.Printf("error=%v app_id=%s member_email=%s", err, c.Get("app_id"), c.Get("member_email"))
	_ = c.JSON(http.StatusInternalServerError, entity.ErrResponse{
		Status:    "INTERNAL_SERVER_ERROR",
		ErrorCode: "UNKNOWN_ERROR",
		Message:   "Sorry, something went wrong",
	})
}
