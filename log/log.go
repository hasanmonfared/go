package log

import (
	"encoding/json"
	"fmt"
	"interface_test/richerror"
	"interface_test/simpleerror"
	"os"
)

type Log struct {
	Errors []richerror.RichError
}

func (l *Log) Append(err error) {
	var finalError richerror.RichError
	rErr, ok := err.(*richerror.RichError)
	if ok {
		finalError = *rErr
	} else {
		sErr, ok := err.(*simpleerror.SimpleError)
		if ok {
			finalError = richerror.RichError{
				Message:   sErr.Output,
				MetaData:  nil,
				Operation: sErr.Operation,
			}
		} else {
			finalError = richerror.RichError{
				Message:   err.Error(),
				MetaData:  nil,
				Operation: "unknown",
			}
		}

	}
	l.Errors = append(l.Errors, finalError)
}
func (l *Log) Save() {
	f, _ := os.OpenFile("errors.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer f.Close()
	data, err := json.Marshal(l.Errors)
	if err != nil {
		fmt.Println("can't marshal data", err)
	}
	f.Write(data)
}
