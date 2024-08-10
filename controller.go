package main

import (
	"KWeb/framework"
	"context"
	"fmt"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	// 这个 channel 负责通知结束
	finish := make(chan struct{}, 1)
	// 这个 channel 负责通知 panic 异常
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	go func() {
		// 增加异常处理
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// Do real action
		time.Sleep(10 * time.Second)
		c.Json(200, "ok")

		finish <- struct{}{}
	}()

	select {
	// 监听 panic
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	// 监听结束事件
	case <-finish:
		fmt.Println("finish")
	// 监听超时事件
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "time out")
		c.SetHasTimeout()
	}
	return nil
}
