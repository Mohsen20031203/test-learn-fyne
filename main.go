package main

import (
    "fmt"
    "github.com/syndtr/goleveldb/leveldb"
)

// تعریف رابط DBClient
type DBClient interface {
    Open() (error, *leveldb.DB)
    Close(db *leveldb.DB)
    Add(key, value string, db leveldb.DB) error
    Get(key string, db leveldb.DB) string
    Read(dbPath string) (error, []dataBace)
}

// ساختار dataBace برای پیاده‌سازی DBClient
type dataBace struct {
    address string
}

func (db *dataBace) Open() (error, *leveldb.DB) {
    // پیاده‌سازی باز کردن دیتابیس
    return nil, nil
}

func (db *dataBace) Close(db *leveldb.DB) {
    // پیاده‌سازی بستن دیتابیس
}

func (db *dataBace) Add(key, value string, db leveldb.DB) error {
    // پیاده‌سازی افزودن داده
    return nil
}

func (db *dataBace) Get(key string, db leveldb.DB) string {
    // پیاده‌سازی دریافت داده
    return ""
}

func (db *dataBace) Read(dbPath string) (error, []dataBace) {
    // پیاده‌سازی خواندن داده‌ها
    return nil, nil
}

// تعریف نوع تابع برای ساخت DBClient
type DBClientFactory func(address string) DBClient

// تابع برای ایجاد نمونه جدید از dataBace
var newDBClient DBClientFactory = func(address string) DBClient {
    return &dataBace{
        address: address,
    }
}

// متغیر جهانی برای نگهداری دیتابیس فعلی
var currentDBClient DBClient

// تابع برای تغییر دیتابیس فعلی
func setDatabase(address string) {
    currentDBClient = newDBClient(address)
}

func main() {
    // انتخاب دیتابیس اول
    setDatabase("localhost:8080")
	currentDBClient.Add()
    fmt.Println("Current database address:", currentDBClient)

    // استفاده از currentDBClient برای عملیات مختلف بر روی دیتابیس

    // تغییر به دیتابیس جدید
    setDatabase("localhost:9090")
    fmt.Println("Current database address:", currentDBClient)
}
