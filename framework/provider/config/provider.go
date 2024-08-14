package config

import (
	"KWeb/framework"
	"KWeb/framework/contract"
)

type KConfigProvider struct {
	c      framework.Container
	folder string
	env    string

	envMaps map[string]string
}

// Register a new function for make a service instance
func (provider *KConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewKConfig
}

// Boot will called when the service instantiate
func (provider *KConfigProvider) Boot(c framework.Container) error {
	provider.folder = c.MustMake(contract.AppKey).(contract.App).ConfigFolder()
	provider.envMaps = c.MustMake(contract.EnvKey).(contract.Env).All()
	provider.env = c.MustMake(contract.EnvKey).(contract.Env).AppEnv()
	provider.c = c
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *KConfigProvider) IsDefer() bool {
	return true
}

// Params define the necessary params for NewInstance
func (provider *KConfigProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.folder, provider.envMaps, provider.env, provider.c}
}

// Name define the name for this service
func (provider *KConfigProvider) Name() string {
	return contract.ConfigKey
}
