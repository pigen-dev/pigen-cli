/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package plugin

import (
	"fmt"

	"github.com/pigen-dev/pigen-cli/internal/plugin"
	"github.com/spf13/cobra"
)

// outputCmd represents the output command
var outputCmd = &cobra.Command{
	Use:   "output",
	Args: cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		pluginName := args[0]
		outputResp := plugin.GetOutput("pigen-plugins.yaml",pluginName)
		if outputResp.Error != nil {
			fmt.Printf("❌ Error getting output: %s\n", outputResp.Error)
		} else {
			// Print the output
			fmt.Println("Output:", outputResp.Output)
			fmt.Println("✅ Output retrieved successfully.")
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// outputCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// outputCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
