package config

import (
	"KWeb/framework"
	"KWeb/framework/contract"
	"path/filepath"
)

type KConfigProvider struct{}

// Register a new function for make a service instance
func (provider *KConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewKConfig
}

// Boot will called when the service instantiate
func (provider *KConfigProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *KConfigProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *KConfigProvider) Params(c framework.Container) []interface{} {
	appService := c.MustMake(contract.AppKey).(contract.App)
	envService := c.MustMake(contract.EnvKey).(contract.Env)
	env := envService.AppEnv()
	configFolder := appService.ConfigFolder()
	envFolder := filepath.Join(configFolder, env)
	return []interface{}{c, envFolder, envService.All()}
}

// Name define the name for this service
func (provider *KConfigProvider) Name() string {
	return contract.ConfigKey
}
