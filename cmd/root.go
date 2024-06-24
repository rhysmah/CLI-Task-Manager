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
	showCompleted   bool
	showIncompleted bool
	tasks           = make(map[string]bool)
)

/*
`rootCmd` is the root command for the CLI application
This is what's called when the CLI application is run.
*/
var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "A simple CLI task manager written in Go.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the CLI Task Manager!")
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your task list.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskDescription := args[0]
		tasks[taskDescription] = false
		fmt.Printf("Adding \"%s\" to your task list...\n", taskDescription)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in your task list.",
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

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your task list as complete.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Marking a task as complete...")
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
	addCmd.Flags().StringP("task", "t", "", "The task to add to your task list.")

	listCmd.Flags().BoolVarP(&showCompleted, "completed", "c", false, "Show completed tasks.")
	listCmd.Flags().BoolVarP(&showIncompleted, "uncompleted", "u", false, "Show uncompleted tasks.")
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
