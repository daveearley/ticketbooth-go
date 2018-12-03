package service

type Error struct {
	StatusCode int
	ErrorMessage error
}

func (e Error) Error() string {
	return e.ErrorMessage.Error()
}
