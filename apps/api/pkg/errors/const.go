package errors

import "fmt"

func ErrMissingField(field string) *ValidationError {
	return &ValidationError{
		ErrorCode: "MISSING_FIELD",
		Message:   fmt.Sprintf("%s field cannot be empty", field),
	}
}

func ErrMissingPathParam(path string) *ValidationError {
	return &ValidationError{
		ErrorCode: "MISSING_PARAMETER",
		Message:   fmt.Sprintf("%s path parameter cannot be empty", path),
	}
}

func ErrInvalidFormat(field, format string) *ValidationError {
	return &ValidationError{
		ErrorCode: "INVALID_FORMAT",
		Message:   fmt.Sprintf("%s field value is not a valid %s", field, format),
	}
}

func ErrInvalidValue(field string) *ValidationError {
	return &ValidationError{
		ErrorCode: "INVALID_VALUE",
		Message:   fmt.Sprintf("%s field value is invalid", field),
	}
}

func ErrUnauthorized(err error, message string) *UnauthorizedError {
	return &UnauthorizedError{
		ErrorCode:     "UNAUTHORIZED",
		Message:       message,
		InternalError: err,
	}
}

func ErrForbidden(err error, message string) *ForbiddenError {
	return &ForbiddenError{
		ErrorCode:     "FORBIDDEN",
		Message:       message,
		InternalError: err,
	}
}

func ErrParseFailed(err error) *ValidationError {
	return &ValidationError{
		ErrorCode:     "PARSE_ERROR",
		Message:       "Failed to parse request body",
		InternalError: err,
	}
}

func ErrInternalServerError(err error, message string) *ServerError {
	return &ServerError{
		ErrorCode:     "INTERNAL_SERVER_ERROR",
		Message:       message,
		InternalError: err,
	}
}

func ErrDataNotFoundError(err error, message string) *NotFoundError {
	return &NotFoundError{
		ErrorCode:     "DATA_NOT_FOUND",
		Message:       message,
		InternalError: err,
	}
}

func DatabaseError(err error, message string) *ServerError {
	return &ServerError{
		ErrorCode:     "DATABASE_ERROR",
		Message:       message,
		InternalError: err,
	}
}

func ErrDuplicateError(err error, message string) *DuplicateError {
	return &DuplicateError{
		ErrorCode:     "DUPLICATE_ERROR",
		Message:       message,
		InternalError: err,
	}
}
