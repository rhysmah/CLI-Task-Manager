package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "\nList all tasks in your task list.",
	Run: func(cmd *cobra.Command, args []string) {

		// Get flag values
		showCompleted, _ := cmd.Flags().GetBool("completed")
		showUncompleted, _ := cmd.Flags().GetBool("uncompleted")

		if len(tasks) == 0 {
			fmt.Println("You currently have no tasks in your list!")
			return
		}

		filteredTasks := filterTasks(tasks, showCompleted, showUncompleted)

		fmt.Println("####################")
		fmt.Println("#    Your Tasks    #")
		fmt.Println("####################")

		index := 1
		for task, isComplete := range filteredTasks {
			if isComplete {
				fmt.Printf("%d. [✓] %s\n", index, task)
			} else {
				fmt.Printf("%d. [ ] %s\n", index, task)
			}
			index++
		}
	},
}

func filterTasks(tasks map[string]bool, showCompleted, showIncompleted bool) map[string]bool {
	if !showCompleted && !showIncompleted {
		return tasks
	}

	filteredTasks := make(map[string]bool)
	for task, isComplete := range tasks {
		if showCompleted && isComplete {
			filteredTasks[task] = isComplete
		} else if showIncompleted && !isComplete {
			filteredTasks[task] = isComplete
		}
	}
	return filteredTasks
}
