package database

import (
	"fmt"
	"log"
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

func AddTask(task string) error {

	return withDatabase(func(db *bolt.DB) error {

		return db.Update(func(t *bolt.Tx) error {
			bucket, err := t.CreateBucketIfNotExists([]byte(bucketName))
			if err != nil {
				return fmt.Errorf("error creating bucket: %w", err)
			}

			falseAsByteSlice := []byte(strconv.FormatBool(false))
			err = bucket.Put([]byte(task), falseAsByteSlice)
			if err != nil {
				return fmt.Errorf("error writing task to database: %w", err)
			}

			log.Printf("Task '%s' successfully written to database.", task)
			return nil
		})
	})
}

func DoTask(task string) error {

	fmt.Println("Accessing tasks...")
	return withDatabase(func(db *bolt.DB) error {

		return db.Update(func(t *bolt.Tx) error {
			bucket := t.Bucket([]byte(bucketName))
			if bucket == nil {
				return fmt.Errorf("bucket '%s' doesn't exist", bucketName)
			}

			taskAsByte := []byte(task)
			cursor := bucket.Cursor()
			k, _ := cursor.Seek(taskAsByte)
			if k == nil || string(k) != task {
				return fmt.Errorf("task '%s' not found", task)
			}

			err := bucket.Put(taskAsByte, []byte(strconv.FormatBool(true)))
			if err != nil {
				return fmt.Errorf("error marking task '%s' as complete: %w", task, err)
			}

			log.Printf("Task '%s' marked as complete", task)
			return nil
		})
	})
}

func ListTasks() (map[string]bool, error) {
	tasksFromDb := make(map[string]bool)

	fmt.Println("Reading tasks in database...")
	err := withDatabase(func(db *bolt.DB) error {

		return db.View(func(t *bolt.Tx) error {
			bucket, err := t.CreateBucketIfNotExists([]byte(bucketName))
			if err != nil {
				return fmt.Errorf("error creating bucket: %w", err)
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

func RemoveTask(task string) error {

	fmt.Println("Attempting to delete task...")
	return withDatabase(func(db *bolt.DB) error {

		return db.Update(func(t *bolt.Tx) error {
			bucket := t.Bucket([]byte(bucketName))
			if bucket == nil {
				return fmt.Errorf("task '%s' not found", task)
			}

			taskAsByte := []byte(task)
			cursor := bucket.Cursor()
			k, _ := cursor.Seek(taskAsByte)
			if k == nil || string(k) != task {
				return fmt.Errorf("task '%s' not found", task)
			}

			err := bucket.Delete(taskAsByte)
			if err != nil {
				return fmt.Errorf("failed to delete task '%s': %w", task, err)
			}

			fmt.Printf("Task '%s' successfully deleted!", task)
			return nil
		})
	})
}

func RemoveAllTasks() error {

	fmt.Println("Attempting to delete all tasks...")

	return withDatabase(func(db *bolt.DB) error {
		return db.Update(func(t *bolt.Tx) error {
			err := t.DeleteBucket([]byte(bucketName))
			if err != nil {
				fmt.Printf("Bucket '%s' not found. Creating bucket...", bucketName)
			}
			fmt.Println("Successfully removed all tasks!")
			t.CreateBucketIfNotExists([]byte(bucketName))
			return nil
		})
	})
}
