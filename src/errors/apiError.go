package errors

type APIError interface {
	HttpStatus() int
	Error() interface{}
}
