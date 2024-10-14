package main

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

// KVData تعریف ساختار داده برای کلید و مقدار
type KVData struct {
	Key   string
	Value string
}

// DBClient interface
type DBClient interface {
	Open() error
	Close()
	Add(key, value string) error
	Get(key string) (string, error)
	Read(start, end *string, count int) (error, []KVData)
	Delete(key string) error
	Iterator(start, end *string) IteratorDB
}

// BadgerClient ساختار برای Badger
type BadgerClient struct {
	db *badger.DB
}

// Open متد برای باز کردن اتصال به پایگاه داده
func (bc *BadgerClient) Open() error {
	var err error
	bc.db, err = badger.Open(badger.DefaultOptions("test"))
	return err
}

// Close متد برای بستن اتصال به پایگاه داده
func (bc *BadgerClient) Close() {
	bc.db.Close()
}

// Add متد برای اضافه کردن کلید و مقدار
func (bc *BadgerClient) Add(key, value string) error {
	return bc.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
}

// Get متد برای دریافت مقدار بر اساس کلید
func (bc *BadgerClient) Get(key string) (string, error) {
	var value string
	err := bc.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		value = string(val)
		return nil
	})
	return value, err
}

// Read متد برای خواندن مقادیر بین start و end
func (bc *BadgerClient) Read(start, end *string, count int) (error, []KVData) {
	var kvData []KVData
	err := bc.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Seek([]byte(*start)); it.Valid(); it.Next() {
			item := it.Item()
			key := string(item.Key())
			val, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}

			kvData = append(kvData, KVData{Key: key, Value: string(val)})

			// اگر تعداد خواندن به حد مشخصی رسید، متوقف می‌شود
			if len(kvData) >= count {
				break
			}

			// اگر به end رسیدیم، متوقف می‌شود
			if end != nil && key > *end {
				break
			}
		}
		return nil
	})
	return err, kvData
}

// Delete متد برای حذف کلید
func (bc *BadgerClient) Delete(key string) error {
	return bc.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

// Iterator متد برای ایجاد یک iterator
func (bc *BadgerClient) Iterator(start, end *string) IteratorDB {
	// این قسمت نیاز به تعریف IteratorDB دارد. به عنوان مثال:
	return &BadgerIterator{
		txn:   bc.db.NewTransaction(false),
		start: start,
		end:   end}
}

// BadgerIterator یک ساختار برای iterator
type BadgerIterator struct {
	txn   *badger.Txn
	start *string
	end   *string
}

func (bi *BadgerIterator) Next() bool {
	// اینجا پیاده‌سازی تابع برای حرکت در iterator
	return false // فقط به عنوان یک مثال
}

func (bi *BadgerIterator) Value() (string, error) {
	// اینجا پیاده‌سازی تابع برای دریافت مقدار
	return "", nil // فقط به عنوان یک مثال
}

// می‌توانید از اینجا شروع کنید
func main() {
	client := &BadgerClient{}

	if err := client.Open(); err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer client.Close()

	// استفاده از متدهای DBClient
	if err := client.Add("key1", "value1"); err != nil {
		fmt.Println("Error adding value:", err)
	}

	val, err := client.Get("key1")
	if err != nil {
		fmt.Println("Error getting value:", err)
	} else {
		fmt.Println("Retrieved value:", val)
	}
}
