package cmd

import (
	"cli-task-manager/database"
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "\nAdd a task to your task list.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		taskDescription := args[0]

		err := database.WriteTask(taskDescription)
		if err != nil {
			fmt.Println("Error writing task to database:", err)
			return
		}
	},
}
