package internal

import (
	"fmt"
	"runtime"
	"strings"
)

var initErrors []string // 用來儲存所有初始化時的錯誤

// 新增初始化時的錯誤
func AddInitError(err error) {
	initErrors = append(initErrors, err.Error())
}

// 取得初始化時的錯誤
func GetInitErrors() string {
	if len(initErrors) == 0 {
		return ""
	}

	return strings.Join(initErrors, "\n") // 用換行符號分隔錯誤訊息
}

// 檢查是否有初始化時的錯誤
func HasInitErrors() bool {
	return len(initErrors) > 0
}

func PrintErrorWithLine(msg string, err error) error {
	msg = msg + ": " + err.Error()

	pc, file, line, ok := runtime.Caller(1)
	if ok {
		fn := runtime.FuncForPC(pc)
		return fmt.Errorf("%s (file: %s, line: %d, func: %s)\n", msg, file, line, fn.Name())
	} else {
		return fmt.Errorf("%s\n", msg)
	}
}