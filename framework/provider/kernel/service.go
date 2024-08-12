package kernel

import (
	"KWeb/framework/gin"
	"net/http"
)

// KKernelService 引擎服务
type KKernelService struct {
	engine *gin.Engine
}

// NewKKernelService 初始化web引擎服务实例
func NewKKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &KKernelService{engine: httpEngine}, nil
}

// HttpEngine 返回web引擎
func (s *KKernelService) HttpEngine() http.Handler {
	return s.engine
}
