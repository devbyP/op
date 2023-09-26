package errorhandler

type HttpError struct {
	Message string
	Code    int
}

func (h HttpError) Error() string {
	return h.Message
}
