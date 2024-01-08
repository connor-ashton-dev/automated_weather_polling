package main

import "fmt"

type Sender interface {
	Send(*WeatherData) error
}

type SMSSender struct {
	number string
}

func NewSMSSender(number string) *SMSSender {
	return &SMSSender{
		number: number,
	}
}
func (s *SMSSender) Send(data *WeatherData) error {
	fmt.Println("sending weather to number:", s.number)
	return nil
}

type EmailSender struct {
	email string
}

func NewEmailSender(email string) *EmailSender {
	return &EmailSender{
		email: email,
	}
}
func (e *EmailSender) Send(data *WeatherData) error {
	fmt.Println("sending weather to email:", e.email)
	return nil
}
