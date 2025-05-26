/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package plugin

import (

	"github.com/spf13/cobra"
)

// pluginCmd represents the plugin command
var PluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func addSubCommands(){
	PluginCmd.AddCommand(listCmd)
	PluginCmd.AddCommand(installCmd)
	PluginCmd.AddCommand(destroyCmd)
	PluginCmd.AddCommand(outputCmd)
	PluginCmd.AddCommand(addCmd)
}

func init() {
	addSubCommands()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pluginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pluginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
