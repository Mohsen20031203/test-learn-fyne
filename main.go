package main

import (
	"fmt"
	"os"
)

func main() {
	// باز کردن یا ایجاد فایل برای نوشتن
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("خطا در باز کردن فایل:", err)
		return
	}
	defer file.Close() // بستن فایل بعد از اتمام کار

	// نوشتن متن در فایل
	text := "سلام، این یک متن نمونه است."
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("خطا در نوشتن در فایل:", err)
		return
	}

	fmt.Println("متن با موفقیت نوشته شد.")
}
