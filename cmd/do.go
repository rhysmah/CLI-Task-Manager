package cmd

import (
	"cli-task-manager/database"
	"fmt"
	"log"
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
			log.Printf("Error marking task '%s' as complete: %v", taskDescription, err)
			fmt.Printf("An error occurred while marking the task as complete: %v\n", err)
			return
		}
		fmt.Printf("Task '%s' successfully marked as complete!\n", taskDescription)
	},
}
