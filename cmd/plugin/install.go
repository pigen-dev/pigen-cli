/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package plugin

import (
	"github.com/pigen-dev/pigen-cli/internal/plugin"
	"github.com/spf13/cobra"
)

var pluginsFile string
// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := plugin.PluginInstall(pluginsFile)
		if err != nil {
			cmd.PrintErr("Error installing plugins: ", err)
		} else {
			cmd.Println("Plugins installed successfully.")
		}
	},
}
func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	installCmd.Flags().StringVarP(&pluginsFile ,"file", "f", "pigen-plugins.yaml", "Your pigen plugins file path")
}
