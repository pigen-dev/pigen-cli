/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package pipeline

import (
	"fmt"

	"github.com/pigen-dev/pigen-cli/internal/pipeline"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate pipeline script",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := pipeline.GenerateScript("pigen-steps.yaml")
		if err != nil {
			fmt.Println("Error generating script:", err)
		} else {
			fmt.Println("Script generated successfully")
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
