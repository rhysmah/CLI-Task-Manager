package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in your task list.",
	Long: `List all tasks in your task list. 
You can use the '--completed' and '--uncompleted' flags to filter by task type.
Using no flags will display both completed and uncompleted tasks.`,
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
		for task, isComplete := range filteredTasks {
			if isComplete {
				fmt.Println("[✓]", task)
			} else {
				fmt.Println("[ ]", task)
			}
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
