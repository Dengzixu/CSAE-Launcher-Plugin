package errorEx

type Error struct {
	Code uint32
	Msg  string
}

func New(code uint32) *Error {
	return &Error{Code: code, Msg: GetMsg(code)}
}

func (err *Error) Error() string {
	return err.Msg
}
