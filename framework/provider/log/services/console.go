package services

import (
	"KWeb/framework"
	"KWeb/framework/contract"
	"os"
)

// KConsoleLog 代表控制台输出
type KConsoleLog struct {
	KLog
}

// NewKConsoleLog 实例化KConsoleLog
func NewKConsoleLog(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	log := &KConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	// 最重要的将内容输出到控制台
	log.SetOutput(os.Stdout)
	log.c = c
	return log, nil
}
