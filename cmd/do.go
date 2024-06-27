package cmd

import (
	"cli-task-manager/database"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "\nMark a task on your task list as complete.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskDescription := strings.Join(args, " ")
		err := database.DoTask(taskDescription)
		if err != nil {
			fmt.Println("Error writing task to database:", err)
			return
		}
	},
}
