package cmd

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
	"todolist/utils"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var viewTaskCommand = &cobra.Command{
	Use:   "view (to be implemented - [-c completed] - view completed tasks | [-uuncompleted] - view uncompleted tasks)",
	Short: "view all tasks",
	Long:  "Use this command followed by task value to include in todo list",
	// Aliases: []string{"viewtasks", "newtask"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := viewTasks(); err != nil {
			return err
		}
		return nil
	},
}

func viewTasks() error {
	db, err := bolt.Open("bolt.db", 0600, &bolt.Options{ReadOnly: true})

	if err != nil {
		return fmt.Errorf("open bolt.db failed: first add a task")
	}

	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket(todoListBucket)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found", todoListBucket)
		}
		var allTaskRows [][]string
		err = bucket.ForEach(func(k, v []byte) error {
			id := binary.BigEndian.Uint64(k)
			pointerToTaskDetail := &taskDetail{}
			json.Unmarshal(v, pointerToTaskDetail)

			keyStr := strconv.FormatUint(id, 10)

			currentRow := []string{
				keyStr,
				string(pointerToTaskDetail.TaskValue),
				humanize.Time(pointerToTaskDetail.CreatedAt),
			}

			allTaskRows = append(allTaskRows, currentRow)
			return nil
		})

		utils.Ytdpretty(allTaskRows)

		if err != nil {
			return fmt.Errorf("Error during tasklist iteration %w", err)
		}
		return nil
	})
	return nil
}

func init() {
	rootCmd.AddCommand(viewTaskCommand)
}
