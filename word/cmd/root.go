package cmd

import "github.com/zhangdapeng520/zdpgo_cobra"

var rootCmd = &zdpgo_cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
}
