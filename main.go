package main

import (
	"fmt"
	"log"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	dbPath := "path/to/your/database"

	// بررسی وجود فایل دیتابیس
	if _, err := os.Stat("/Users/macbookpro/Desktop/helo"); os.IsNotExist(err) {
		// اگر دیتابیس وجود ندارد، آن را ایجاد کنید
		db, err := leveldb.OpenFile(dbPath, nil)
		if err != nil {
			log.Fatalf("Failed to create database: %v", err)
		}
		defer db.Close()

		fmt.Println("Database created successfully.")
	} else {
		// اگر دیتابیس وجود دارد، کار دیگری انجام دهید
		fmt.Println("Database already exists. Performing other actions.")

		// کد مورد نیاز برای انجام کار دیگر
		// ...
	}
}
