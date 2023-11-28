package errors

type BadRequestError struct {
	Code          int
	Status        string
	Message       string
	InternalError error
}

func (e BadRequestError) Error() string {
	return e.Message
}

type ValidationError struct {
	Code          int
	Status        string
	Message       string
	InternalError error
}

func (e ValidationError) Error() string {
	return e.Message
}

type ServerError struct {
	Code          int
	Status        string
	Message       string
	InternalError error
}

func (e ServerError) Error() string {
	return e.Message
}

type UnauthorizedError struct {
	Code          int
	Status        string
	Message       string
	InternalError error
}

func (e UnauthorizedError) Error() string {
	return e.Message
}

type ForbiddenError struct {
	Code          int
	Status        string
	Message       string
	InternalError error
}

func (e ForbiddenError) Error() string {
	return e.Message
}

type NotFoundError struct {
	Code          int
	Status        string
	Message       string
	InternalError error
}

func (e NotFoundError) Error() string {
	return e.Message
}

type DuplicateError struct {
	Code          int
	Status        string
	Message       string
	InternalError error
}

func (e DuplicateError) Error() string {
	return e.Message
}
