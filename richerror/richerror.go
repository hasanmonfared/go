package richerror

type RichError struct {
	Message   string
	MetaData  map[string]string
	Operation string
}

func (r RichError) Error() string {
	return r.Message
}
