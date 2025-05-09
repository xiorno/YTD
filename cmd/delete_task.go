package cmd

import (
	"fmt"
	"strconv"
	"todolist/utils"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var deleteTaskCommand = &cobra.Command{
	Use:          "delete task-id",
	Short:        "delete a task",
	SilenceUsage: true,
	Long:         "Use this command followed by task value to include in todo list",
	// Aliases: []string{"viewtasks", "newtask"},
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskID, err := strconv.ParseUint(args[0], 10, 64)

		if err != nil {
			return fmt.Errorf("Invalid task Id: %s, provide a numical id", args[0])
		}
		if err := deleteTask(taskID); err != nil {
			return err
		}
		return nil
	},
}

func deleteTask(taskID uint64) error {
	db, err := bolt.Open("bolt.db", 0600, nil)
	// error handling for opening databse
	if err != nil {
		return fmt.Errorf("error occured when opening db")
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(todoListBucket)
		// error handling for opening bucket
		if bucket == nil {
			return fmt.Errorf("Bucket %q does not exist or is not initialized. No tasks to delete.", todoListBucket)
		}

		deleteKey := utils.IntToByte(taskID)
		bucket.Delete(deleteKey)
		return nil
	})

	if err != nil {
		return fmt.Errorf("transaction failed during deletion")
	}
	fmt.Printf("Task with ID %v deleted succesfully\n", taskID)
	return nil
}
func init() {
	rootCmd.AddCommand(deleteTaskCommand)
}
