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

			log.Printf("Successfully added task '%s' to %s", task, bucketName)
			return nil
		})
	})
}

func DoTask(task string) error {
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

			log.Printf("Successfully completed task '%s' in %s", task, bucketName)
			return nil
		})
	})
}

func ListTasks() (map[string]bool, error) {
	tasksFromDb := make(map[string]bool)

	err := withDatabase(func(db *bolt.DB) error {
		return db.View(func(t *bolt.Tx) error {
			bucket := t.Bucket([]byte(bucketName))
			if bucket == nil {
				return fmt.Errorf("bucket '%s' doesn't exist", bucketName)
			}

			return bucket.ForEach(func(k, v []byte) error {
				taskAsString := string(k)
				valueAsBool, err := strconv.ParseBool(string(v))
				if err != nil {
					return fmt.Errorf("error retrieving task %s: %w", taskAsString, err)
				}

				tasksFromDb[taskAsString] = valueAsBool
				return nil
			})
		})
	})

	if err != nil {
		return nil, err
	}

	log.Printf("Successfully retrieved all tasks from %s", bucketName)
	return tasksFromDb, nil
}

func RemoveTask(task string) error {
	return withDatabase(func(db *bolt.DB) error {

		return db.Update(func(t *bolt.Tx) error {
			bucket := t.Bucket([]byte(bucketName))
			if bucket == nil {
				return fmt.Errorf("bucket '%s' not found", bucketName)
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

			log.Printf("Successfully removed task '%s' from %s", task, bucketName)
			return nil
		})
	})
}

func RemoveAllTasks() error {
	return withDatabase(func(db *bolt.DB) error {

		return db.Update(func(t *bolt.Tx) error {
			err := t.DeleteBucket([]byte(bucketName))
			if err != nil && err != bolt.ErrBucketNotFound {
				return fmt.Errorf("failed to delete bucket '%s': %w", bucketName, err)
			}

			if err == bolt.ErrBucketNotFound {
				log.Printf("Bucket '%s' not found. Creating a new bucket...", bucketName)
			} else {
				log.Printf("Bucket '%s' successfully removed. Creating a new empty bucket...", bucketName)
			}

			_, err = t.CreateBucket([]byte(bucketName))
			if err != nil {
				return fmt.Errorf("failed to create bucket '%s': %w", bucketName, err)
			}

			log.Printf("Successfully removed all tasks from %s", bucketName)
			return nil
		})
	})
}
