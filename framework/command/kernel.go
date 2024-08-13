package command

import "KWeb/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	// app
	root.AddCommand(initAppCommand())
	// cron
	root.AddCommand(initCronCommand())
}
