package apperr

import (
	"fmt"
	"strings"

	"github.com/rhizomplatform/golib/logger"
)

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

func NewErr(opt Options) *AppErr {
	appErr := &AppErr{
		HTTPCode: opt.HTTPCode,
		Err:      opt.Err,
		Key:      opt.Key,
		Message:  opt.Message,
		Data:     opt.Data,
	}
	if opt.DoNotLogErr {
		return appErr
	}
	logger.ErrorApp(appErr.Error())
	return appErr
}

func (appErr AppErr) Error() string {
	var elements = make([]string, 0, 5)
	if appErr.HTTPCode > 0 {
		elements = append(elements, fmt.Sprintf("http: %d", appErr.HTTPCode))
	}
	if appErr.Err != nil {
		elements = append(elements, fmt.Sprintf("err: %v", appErr.Err))
	}
	if appErr.Key != "" {
		elements = append(elements, fmt.Sprintf("key: %s", appErr.Key))
	}
	if appErr.Message != "" {
		elements = append(elements, fmt.Sprintf("msg: %s", appErr.Message))
	}
	return strings.Join(elements, ", ")
}
