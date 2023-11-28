package errors

import "fmt"

func ErrMissingField(field string) *ValidationError {
	return &ValidationError{
		Code:    400,
		Status:  "MISSING_FIELD",
		Message: fmt.Sprintf("%s field cannot be empty", field),
	}
}

func ErrMissingPathParam(path string) *ValidationError {
	return &ValidationError{
		Code:    400,
		Status:  "MISSING_PARAMETER",
		Message: fmt.Sprintf("%s path parameter cannot be empty", path),
	}
}

func ErrInvalidFormat(field, format string) *ValidationError {
	return &ValidationError{
		Code:    400,
		Status:  "INVALID_FORMAT",
		Message: fmt.Sprintf("%s field value is not a valid %s", field, format),
	}
}

func ErrInvalidValue(field string) *ValidationError {
	return &ValidationError{
		Code:    400,
		Status:  "INVALID_VALUE",
		Message: fmt.Sprintf("%s field value is invalid", field),
	}
}

func ErrUnauthorized(err error, message string) *UnauthorizedError {
	return &UnauthorizedError{
		Code:          401,
		Status:        "UNAUTHORIZED",
		Message:       message,
		InternalError: err,
	}
}

func ErrForbidden(err error, message string) *ForbiddenError {
	return &ForbiddenError{
		Code:          403,
		Status:        "FORBIDDEN",
		Message:       message,
		InternalError: err,
	}
}

func ErrParseFailed(err error) *ValidationError {
	return &ValidationError{
		Code:          400,
		Status:        "PARSE_ERROR",
		Message:       "Failed to parse request body",
		InternalError: err,
	}
}

func ErrInternalServerError(err error, message string) *ServerError {
	return &ServerError{
		Code:          500,
		Status:        "INTERNAL_SERVER_ERROR",
		Message:       message,
		InternalError: err,
	}
}

func ErrDataNotFoundError(err error, message string) *NotFoundError {
	return &NotFoundError{
		Code:          404,
		Status:        "DATA_NOT_FOUND",
		Message:       message,
		InternalError: err,
	}
}

func DatabaseError(err error, message string) *ServerError {
	return &ServerError{
		Code:          500,
		Status:        "DATABASE_ERROR",
		Message:       message,
		InternalError: err,
	}
}

func ErrDuplicateError(err error, message string) *DuplicateError {
	return &DuplicateError{
		Code:          409,
		Status:        "DUPLICATE_ERROR",
		Message:       message,
		InternalError: err,
	}
}
