package database

import (
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
)

const dbFileName = "tasks.db"
const bucketName = "Tasks"
const readWriteCode = 0600

func withDatabase(operation func(db *bolt.DB) error) error {

	fmt.Println("Accessing database...")
	db, err := bolt.Open(dbFileName, readWriteCode, nil)
	if err != nil {
		return fmt.Errorf("error opening database: %s", err)
	}

	defer db.Close()

	err = operation(db)
	if err != nil {
		return err
	}

	return nil // Operations successful
}

func ListTasks() (map[string]bool, error) {

	tasksFromDb := make(map[string]bool)

	fmt.Println("Reading tasks in database...")
	err := withDatabase(func(db *bolt.DB) error {

		return db.View(func(t *bolt.Tx) error {

			// Attempt to find bucket containing tasks
			bucket := t.Bucket([]byte(bucketName))
			if bucket == nil {
				return fmt.Errorf("bucket %s not found: ", bucketName)
			}

			// Bucket found; iterate through tasks; write to map
			return bucket.ForEach(func(k, v []byte) error {
				taskAsString := string(k)

				valueAsBool, err := strconv.ParseBool(string(v))
				if err != nil {
					return fmt.Errorf("error retrieving tasks %s", taskAsString)
				}

				tasksFromDb[taskAsString] = valueAsBool
				return nil // operation successful
			})
		})
	})

	// Error performing operation
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully read tasks!")
	return tasksFromDb, nil
}

func WriteTask(task string) error {

	return withDatabase(func(db *bolt.DB) error {
		err := db.Update(func(transaction *bolt.Tx) error {

			bucket, err := transaction.CreateBucketIfNotExists([]byte(bucketName))
			if err != nil {
				return fmt.Errorf("error creating bucket: %s", err)
			}

			// Default value for all tasks is Boolean false
			falseAsByteSlice := []byte(strconv.FormatBool(false))
			err = bucket.Put([]byte(task), falseAsByteSlice)
			if err != nil {
				return fmt.Errorf("error writing task to database: %s", err)
			}
			return nil // Task successfull written to database
		})

		if err != nil {
			return err
		}

		fmt.Printf("Task '%s' successfully written to database!\n", task)
		return nil
	})
}
