package command

import "KWeb/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	root.AddCommand(initAppCommand())
}
