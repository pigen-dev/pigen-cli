/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package pipeline

import (
	"fmt"

	"github.com/pigen-dev/pigen-cli/internal/pipeline"
	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "link github repo and create trigger",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setup called")
		err := pipeline.SetupPipeline(pigenStepsPath)
		if err != nil {
			fmt.Printf("Error setting up pipeline: %v\n", err)
		} else {
			fmt.Println("✅ Pipeline setup successfully")
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
