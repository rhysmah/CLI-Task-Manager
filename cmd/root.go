/*
Copyright © 2024 Rhys Mahannah <https://github.com/rhysmah>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	tasks = make(map[string]bool)
)

/*
`rootCmd` is the root command for the CLI application
This is what's called when the CLI application is run.
*/
var rootCmd = &cobra.Command{
	Use:   "task [command]",
	Short: "A simple CLI task manager written in Go.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the CLI Task Manager!")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	tasks = map[string]bool{
		"task1": false,
		"task2": true,
	}

	/*
		Add commands to the root command. This adds actions to the CLI
		application. When the user runs the CLI application, they can
		choose to run one of these commands, which will execute their
		respective functions defined in the `Run` field.
	*/
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(doCmd)
	rootCmd.AddCommand(listCmd)

	/*
		Add flags to the commands. Flags are options that the user can
		pass to the command when they run the CLI application.
	*/
	listCmd.Flags().BoolP("completed", "c", false, "Show completed tasks.")
	listCmd.Flags().BoolP("uncompleted", "u", false, "Show uncompleted tasks.")
}
