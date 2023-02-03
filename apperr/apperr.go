package apperr

import (
	"errors"
	"fmt"
	"strconv"
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
	if opt.NotPrint {
		return appErr
	}
	logger.ErrorApp(appErr.Error())
	return appErr
}

func (appErr AppErr) Error() string {
	var elements = make([]string, 0, 4)
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

func GetFromString(str string) *AppErr {
	var (
		elements = strings.Split(str, ", ")
		appERR   = &AppErr{}
	)
	if len(elements) == 0 {
		return nil
	}
	for _, element := range elements {
		switch {
		case strings.HasPrefix(element, "http: "):
			appERR.HTTPCode, _ = strconv.Atoi(strings.TrimPrefix(element, "http: "))
		case strings.HasPrefix(element, "err: "):
			appERR.Err = errors.New(strings.TrimPrefix(element, "err: "))
		case strings.HasPrefix(element, "key: "):
			appERR.Key = strings.TrimPrefix(element, "key: ")
		case strings.HasPrefix(element, "msg: "):
			appERR.Message = strings.TrimPrefix(element, "msg: ")
		}
	}
	return appERR
}
