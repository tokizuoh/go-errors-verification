package myerror

func New(text string) error {
	return errorString(text)
}

type errorString string

func (e errorString) Error() string {
	return string(e)
}
