package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (u *User) Save() {
	fmt.Println("User saved!")
}

type EmailService struct{}

func (e *EmailService) SendEmail(u *User) {
	fmt.Println("Email sent to", u.Email)
}

func main() {
	u := &User{Name: "Alice", Email: "alice@example.com"}
	u.Save()

	emailService := &EmailService{}
	emailService.SendEmail(u)
}
