package env

import (
	"KWeb/framework"
	"KWeb/framework/contract"
)

type KEnvProvider struct {
	Folder string
}

// Register a new function for make a service instance
func (provider *KEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewKEnv
}

// Boot will called when the service instantiate
func (provider *KEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *KEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *KEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.Folder}
}

// / Name define the name for this service
func (provider *KEnvProvider) Name() string {
	return contract.EnvKey
}
