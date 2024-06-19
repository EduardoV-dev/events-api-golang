package utils

type HttpError struct {
	Message string
	Status  int
  Error error
}

func NewHttpError(err error, status int) *HttpError {
	return &HttpError{
    Error: err,
    Message: err.Error(),
    Status: status,
	}
}
