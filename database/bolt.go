package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

const dbFileName = "tasks.db"
const bucketName = "Tasks"
const readWriteCode = 0600

func WriteTask(task string) error {

	fmt.Println("Opening database...")
	db, err := bolt.Open(dbFileName, readWriteCode, nil)
	if err != nil {
		return fmt.Errorf("error opening database: %s", err)
	}

	defer db.Close()

	// Begin a read-write transaction
	fmt.Println("Beginning to write task to database...")
	err = db.Update(func(transaction *bolt.Tx) error {

		// Creates bucket if it doesn't exist; a bucket is a collection
		// of key-value pairs -- tasks and their status (complete or incomplete)
		bucket, err := transaction.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("error creating bucket: %s", err)
		}

		// Default value for all task keys is "false"
		// `bucket` requires the key and value to both be bytes
		stringVal := strconv.FormatBool(false)
		falseByteSlice := []byte(stringVal)

		err = bucket.Put([]byte(task), falseByteSlice)
		if err != nil {
			return fmt.Errorf("error writing task to database: %s", err)
		}
		return nil
	})

	if err != nil {
		log.Printf("Error writing task '%s' to database: %s\n", task, err)
		return err
	}

	fmt.Printf("Task '%s' written to database successfully!\n", task)
	return nil
}
