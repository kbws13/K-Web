package services

import (
	"KWeb/framework"
	"KWeb/framework/contract"
	"io"
)

type KCustomLog struct {
	KLog
}

func NewKCustomLog(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)
	output := params[4].(io.Writer)

	log := &KConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)
	log.SetOutput(output)
	log.c = c
	return log, nil
}
