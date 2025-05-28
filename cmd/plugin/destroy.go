/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package plugin

import (
	"fmt"
	"github.com/pigen-dev/pigen-cli/internal/plugin"
	"github.com/spf13/cobra"
)

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Args: cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		pluginName := args[0]
		// 1. Destroy logic here
		err := plugin.DestroyPlugin("pigen-plugins.yaml", pluginName)
		if err != nil {
			fmt.Printf("❌ Failed to destroy plugin: %s error: %s\n", pluginName, err)
			return
		}
		fmt.Printf("✅ Plugin \"%s\" destroyed successfully.\n", pluginName)

		// 2. Get the value and check if the flag was explicitly passed
		updateYaml, _ := cmd.Flags().GetBool("update-yaml")

		if !cmd.Flags().Changed("update-yaml") {
				// Ask the user if they want to update plugin.yaml
				fmt.Printf("❓ Do you want to remove \"%s\" from plugin.yaml as well? [y/N]: ", pluginName)

				var response string
				fmt.Scanln(&response)

				if response == "y" || response == "Y" {
						updateYaml = true
				} else {
						updateYaml = false
				}
		}
		if updateYaml {
			// 3. Update the plugin.yaml file
			err := plugin.UpdatePluginYaml("pigen-plugins.yaml", pluginName)
			if err != nil {
				fmt.Printf("❌ Failed to update plugin.yaml: %s\n", err)
				return
			}
			fmt.Printf("✅ Plugin \"%s\" removed from plugin.yaml successfully.\n", pluginName)
		} else {
			fmt.Println("❌ Plugin not removed from plugin.yaml.")
		}
		// 4. Print the final message
		fmt.Println("✅ Plugin destroyed successfully.")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// destroyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	destroyCmd.Flags().Bool("update-yaml", false, "Update plugins file after destroy")
}
