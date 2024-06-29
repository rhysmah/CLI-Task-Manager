package cmd

import (
	"cli-task-manager/database"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "\nAdd a task to your task list.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		taskDescription := strings.Join(args, " ")
		err := database.AddTask(taskDescription)
		if err != nil {
			fmt.Printf("error writing task '%s': %v\n", taskDescription, err)
			return
		}
		fmt.Printf("Task '%s' successfully written to database!\n", taskDescription)
	},
}
