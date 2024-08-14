package command

import (
	"KWeb/framework/cobra"
	"KWeb/framework/contract"
	"fmt"
)

// envCommand show current environment
var envCommand = &cobra.Command{
	Use:   "env",
	Short: "get current environment",
	Run: func(c *cobra.Command, args []string) {
		container := c.GetContainer()
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		fmt.Println("environment:", envService.AppEnv())
	},
}
