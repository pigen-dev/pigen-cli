/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package pipeline

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pigen-steps.yaml path
var pigenStepsPath string
// pipelineCmd represents the pipeline command
var PipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pipeline called")
	},
}

func addSubCommands(){
	PipelineCmd.AddCommand(generateCmd)
	PipelineCmd.AddCommand(setupCmd)
}

func init() {
	addSubCommands()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	PipelineCmd.PersistentFlags().StringVarP(&pigenStepsPath, "file", "f", "pigen-steps.yaml", "pigen-steps.yaml file path")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pipelineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
