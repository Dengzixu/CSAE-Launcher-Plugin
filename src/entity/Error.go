package entity

const (
	InternalError = 500
)

type Errors struct {
	Code    int
	Message string
	Error   error
}
