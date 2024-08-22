package trace

import (
	"KWeb/framework"
	"KWeb/framework/contract"
)

type KTraceProvider struct {
	c framework.Container
}

// Register registe a new function for make a service instance
func (provider *KTraceProvider) Register(c framework.Container) framework.NewInstance {
	return NewKTraceService
}

// Boot will called when the service instantiate
func (provider *KTraceProvider) Boot(c framework.Container) error {
	provider.c = c
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *KTraceProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *KTraceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.c}
}

// / Name define the name for this service
func (provider *KTraceProvider) Name() string {
	return contract.TraceKey
}
