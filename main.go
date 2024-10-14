package main

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

func main() {
	opt := badger.DefaultOptions("test")
	db, err := badger.Open(opt)
	fmt.Println("true")
	if err != nil {
		return
	}
	defer db.Close()

	// اضافه کردن 100 کلید/مقدار به دیتابیس
	err = db.Update(func(txn *badger.Txn) error {
		for i := 0; i < 100; i++ {
			key := fmt.Sprintf("key%d", i)
			value := fmt.Sprintf("value%d", i)
			err := txn.Set([]byte(key), []byte(value))
			if err != nil {
				return err
			}
			fmt.Println("Inserted:", key)
		}
		return nil
	})
	if err != nil {
		return
	}

	// پیمایش (iteration) داده‌ها از دیتابیس
	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 1 // تنظیم میزان پیش‌بارگذاری (prefetching)
		it := txn.NewIterator(opts)
		defer it.Close()

		fmt.Println("Iterating through keys and values:")
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				fmt.Printf("Key: %s, Value: %s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return
	}
}
