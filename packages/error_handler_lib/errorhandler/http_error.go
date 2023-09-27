package errorhandler

type HttpError struct {
	Message string
	Code    int
}

func NewHttpError(err string, code int) HttpError {
	return HttpError{
		Message: err,
		Code:    code,
	}
}

func (h HttpError) Error() string {
	return h.Message
}
