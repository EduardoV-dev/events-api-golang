package utils

type HttpError struct {
	Message string
	Status  int
  Err error
}

func NewHttpError(err error, status int) *HttpError {
  if err == nil {
    return nil
  }
  
	return &HttpError{
    Err: err,
    Message: err.Error(),
    Status: status,
	}
}
