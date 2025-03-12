// Package magazine
package magazine

import "fmt"

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	Address
}

type Employee struct {
	Name        string
	Salary      float64
	Address
}

func PrintInfo(s *Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly rate", s.Rate)
	fmt.Println("Active?", s.Active)
}

func DefaultSubscriber(name string) *Subscriber {
	var s Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}

func ApplyDiscount(s *Subscriber) {
	s.Rate = 4.99 // 使用点运算符在struct指针和struct上都可访问字段
}
