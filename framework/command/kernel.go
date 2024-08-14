package command

import "KWeb/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	// app
	root.AddCommand(initAppCommand())
	// env
	root.AddCommand(envCommand)
	// cron
	//root.AddCommand(initCronCommand())
}
