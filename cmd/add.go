package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "\nAdd a task to your task list.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskDescription := args[0]
		tasks[taskDescription] = false
		fmt.Printf("Added \"%s\" to your task list...\n", taskDescription)
	},
}
