/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli-task-manager/database"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a task from your task list",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		removeAll, _ := cmd.Flags().GetBool("all")
		if removeAll {
			err := database.RemoveAllTasks()
			if err != nil {
				log.Printf("Error removing all tasks: %v", err)
				return
			}
			log.Println("Successfully removed all tasks")
			return
		}

		// Remove specific task
		taskDescription := strings.Join(args, " ")
		err := database.RemoveTask(taskDescription)
		if err != nil {
			log.Printf("Error removing task '%s': %v", taskDescription, err)
			return
		}
		log.Printf("Task '%s' successfully removed", taskDescription)
	},
}
