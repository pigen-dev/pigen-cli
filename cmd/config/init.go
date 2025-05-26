/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"fmt"

	"os"

	"github.com/pigen-dev/pigen-cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := config.Config{}
		// Bubble Tea: Choose cloud provider
		menu := menuModel{options: []string{"GCP", "AWS", "AZURE"}}
		p := tea.NewProgram(menu)
		finalModel, err := p.Run()
		if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
		}

		choice := finalModel.(menuModel).choice
		fmt.Println("You selected:", choice)

		if choice == "GCP" {
				config.PigenCore.CloudProvider.Type = "GCP"
				// Bubble Tea: Prompt for project_id and region
				fmt.Print("Enter your GCP Project ID: ")
				projectInput := textinput.New()
				projectInput.Placeholder = "Project ID"
				projectInput.Focus()
				projectInput.CharLimit = 64
				projectInput.Width = 30

				regionInput := textinput.New()
				regionInput.Placeholder = "Region"
				regionInput.CharLimit = 64
				regionInput.Width = 30

				form := formModel{
						inputs: []textinput.Model{projectInput, regionInput},
						focus:  0,
				}

				formProgram := tea.NewProgram(form)
				finalForm, err := formProgram.Run()
				if err != nil {
						fmt.Println("Error:", err)
						os.Exit(1)
				}

				f := finalForm.(formModel)
				projectID := f.inputs[0].Value()
				region := f.inputs[1].Value()

				fmt.Println("\n--- GCP Configuration ---")
				fmt.Println("Project ID:", projectID)
				fmt.Println("Region:", region)
				config.PigenCore.CloudProvider.ProjectID = projectID
				config.PigenCore.CloudProvider.Region = region
		}
		err = config.InitConfig()
		if err != nil {
				fmt.Println("Error initializing config:", err)
				os.Exit(1)
		}
		fmt.Println("Configuration initialized successfully.")
		fmt.Println("Config file created at:", viper.ConfigFileUsed())
		fmt.Println("You can now use 'pigen' cli tool")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


// --- Menu Model ---
type menuModel struct {
	options []string
	selected int
	done bool
	choice string
}

func (m menuModel) Init() tea.Cmd { return nil }

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
			switch msg.String() {
			case "up":
					if m.selected > 0 {
							m.selected--
					}
			case "down":
					if m.selected < len(m.options)-1 {
							m.selected++
					}
			case "enter":
					m.done = true
					m.choice = m.options[m.selected]
					return m, tea.Quit
			case "q":
					return m, tea.Quit
			}
	}
	return m, nil
}

func (m menuModel) View() string {
	if m.done {
			return ""
	}
	s := "Select Cloud Provider:\n\n"
	for i, option := range m.options {
			cursor := " "
			if i == m.selected {
					cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, option)
	}
	s += "\n↑/↓ to move, enter to select"
	return s
}

// --- Form Model for GCP ---
type formModel struct {
	inputs []textinput.Model
	focus  int
	done   bool
}

func (m formModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m formModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
			switch msg.String() {
			case "enter":
					if m.focus == len(m.inputs)-1 {
							m.done = true
							return m, tea.Quit
					}
					m.focus++
			case "tab", "down":
					m.focus = (m.focus + 1) % len(m.inputs)
			case "up":
					m.focus = (m.focus - 1 + len(m.inputs)) % len(m.inputs)
			}

	}

	for i := range m.inputs {
			if i == m.focus {
					m.inputs[i].Focus()
			} else {
					m.inputs[i].Blur()
			}
			m.inputs[i], _ = m.inputs[i].Update(msg)
	}

	return m, nil
}

func (m formModel) View() string {
	if m.done {
			return "Form submitted successfully.\n"
	}

	s := "Enter GCP Configuration:\n\n"
	for _, input := range m.inputs {
			s += input.View() + "\n"
	}
	s += "\nTab to switch, enter to submit."
	return s
}
