package main

import (
	"os"
	"testing"
)

func TestFileContent(t *testing.T) {
	// خواندن محتوای فایل
	content, err := os.ReadFile("example.txt")
	if err != nil {
		t.Fatalf("خطا در خواندن فایل: %v", err)
	}

	// متن مورد انتظار
	expectedText := ""

	// بررسی اینکه محتوای فایل با متن مورد انتظار برابر است یا خیر
	if string(content) == expectedText {
		t.Errorf("متن فایل اشتباه است. انتظار داشتیم: %s\n ولی یافتیم: %s", expectedText, string(content))
	} else {
		t.Log("تست با موفقیت انجام شد. متن صحیح است.")
	}
}

func TestMn(t *testing.T) {
	content, err := os.ReadFile("example.txt")
	if err != nil {
		t.Fatalf("خطا در خواندن فایل: %v", err)
	}

	if content != nil {
		t.Log("yes")
	}
}
