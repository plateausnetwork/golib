package apperr

import (
	"fmt"
	"strings"

	"github.com/rhizomplatform/golib/logger"
)

type AppErr struct {
	HTTPCode int
	Err      error
	Key      string
	Message  string
}

func New(httpCode int, err error, msg, key string) *AppErr {
	appErr := &AppErr{
		HTTPCode: httpCode,
		Err:      err,
		Key:      key,
		Message:  msg,
	}
	logger.ErrorApp(appErr.Error())
	return appErr
}

func (appErr AppErr) Error() string {
	var messages []string
	if appErr.HTTPCode > 0 {
		messages = append(messages, fmt.Sprintf("http: %d", appErr.HTTPCode))
	}
	if appErr.Err != nil {
		messages = append(messages, fmt.Sprintf("err: %v", appErr.Err))
	}
	if appErr.Key != "" {
		messages = append(messages, fmt.Sprintf("key: %s", appErr.Key))
	}
	if appErr.Message != "" {
		messages = append(messages, fmt.Sprintf("msg: %s", appErr.Message))
	}
	return strings.Join(messages, ", ")
}
