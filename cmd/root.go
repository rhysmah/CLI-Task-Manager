/*
Copyright © 2024 Rhys Mahannah <https://github.com/rhysmah>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(doCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(removeCmd)

	listCmd.Flags().BoolP("completed", "c", false, "Show completed tasks.")
	listCmd.Flags().BoolP("uncompleted", "u", false, "Show uncompleted tasks.")
	removeCmd.Flags().BoolP("all", "a", false, "Remove all tasks.")
}
