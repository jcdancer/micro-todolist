package handlers

import (
	"api-gateway/pkg/logging"
	"errors"
)

// PanicIfUserError
// @Description: 包装错误
// @param err
func PanicIfUserError(err error) {
	if err != nil {
		err := errors.New("userService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}

func PanicIfTaskError(err error) {
	if err != nil {
		err := errors.New("taskService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}