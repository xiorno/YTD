package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"todolist/utils"

	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

type taskDetail struct {
	TaskValue []byte
	CreatedAt time.Time
}

var addTaskCommand = &cobra.Command{
	Use:     "add taskname",
	Short:   "add a task to list",
	Long:    "Use this command followed by task value to include in todo list",
	Aliases: []string{"addtask", "newtask"},
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := addTask(args[0]); err != nil {
			return err
		}
		return nil
	},
}

var todoListBucket = []byte("TodoList")

func (td taskDetail) toByte() []byte {
	tdByte, _ := json.Marshal(td)
	return tdByte
}

func addTask(task string) error {
	switch {
	case task == "":
		return fmt.Errorf("task is an empty string")
	case len(task) <= 4:
		return fmt.Errorf("Task name is too short (min. 4 characters)")
	}

	db, err := bolt.Open("bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(todoListBucket)
		if err != nil {
			return err
		}
		id, _ := bucket.NextSequence()
		key := utils.IntToByte(id)
		value := taskDetail{
			TaskValue: []byte(task),
			CreatedAt: time.Now(),
		}

		err = bucket.Put(key, value.toByte())
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}
	fmt.Printf("added: %q\n", task)
	return nil
}

func init() {
	rootCmd.AddCommand(addTaskCommand)
}
