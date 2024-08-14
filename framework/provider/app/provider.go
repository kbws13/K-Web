package app

import (
	"KWeb/framework"
	"KWeb/framework/contract"
)

// KAppProvider 提供App的具体实现方法
type KAppProvider struct {
	BaseFolder string
}

// Register 注册KApp方法
func (h *KAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewKApp
}

// Boot 启动调用
func (h *KAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *KAppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (h *KAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, h.BaseFolder}
}

// Name 获取字符串凭证
func (h *KAppProvider) Name() string {
	return contract.AppKey
}
