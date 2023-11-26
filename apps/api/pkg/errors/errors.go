package errors

type BadRequestError struct {
	ErrorCode     string
	Message       string
	InternalError string
}

func (e BadRequestError) Error() string {
	return e.Message
}

type ValidationError struct {
	ErrorCode     string
	Message       string
	InternalError error
}

func (e ValidationError) Error() string {
	return e.Message
}

type ServerError struct {
	ErrorCode     string
	Message       string
	InternalError error
}

func (e ServerError) Error() string {
	return e.Message
}

type UnauthorizedError struct {
	ErrorCode     string
	Message       string
	InternalError error
}

func (e UnauthorizedError) Error() string {
	return e.Message
}

type ForbiddenError struct {
	ErrorCode     string
	Message       string
	InternalError error
}

func (e ForbiddenError) Error() string {
	return e.Message
}

type NotFoundError struct {
	ErrorCode     string
	Message       string
	InternalError error
}

func (e NotFoundError) Error() string {
	return e.Message
}

type DuplicateError struct {
	ErrorCode     string
	Message       string
	InternalError error
}

func (e DuplicateError) Error() string {
	return e.Message
}
