package database

import (
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
)

const (
	dbFileName    = "tasks.db"
	bucketName    = "Tasks"
	readWriteCode = 0600
)

func withDatabase(operation func(db *bolt.DB) error) error {
	fmt.Println("Accessing database...")
	db, err := bolt.Open(dbFileName, readWriteCode, nil)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}
	defer db.Close()

	return operation(db)
}

func ListTasks() (map[string]bool, error) {
	tasksFromDb := make(map[string]bool)

	fmt.Println("Reading tasks in database...")
	err := withDatabase(func(db *bolt.DB) error {

		return db.View(func(t *bolt.Tx) error {
			bucket := t.Bucket([]byte(bucketName))
			if bucket == nil {
				return fmt.Errorf("bucket %s not found: ", bucketName)
			}

			return bucket.ForEach(func(k, v []byte) error {
				taskAsString := string(k)
				valueAsBool, err := strconv.ParseBool(string(v))
				if err != nil {
					return fmt.Errorf("error retrieving tasks %s: %w", taskAsString, err)
				}

				tasksFromDb[taskAsString] = valueAsBool
				return nil
			})
		})
	})

	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully read tasks!")
	return tasksFromDb, nil
}

func WriteTask(task string) error {
	return withDatabase(func(db *bolt.DB) error {
		return db.Update(func(transaction *bolt.Tx) error {

			bucket, err := transaction.CreateBucketIfNotExists([]byte(bucketName))
			if err != nil {
				return fmt.Errorf("error creating bucket: %w", err)
			}

			falseAsByteSlice := []byte(strconv.FormatBool(false))
			err = bucket.Put([]byte(task), falseAsByteSlice)
			if err != nil {
				return fmt.Errorf("error writing task to database: %w", err)
			}
			fmt.Printf("Task '%s' successfully written to database!\n", task)
			return nil
		})
	})
}
