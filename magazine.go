// Package magazine
package magazine

import "fmt"

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func PrintInfo(s *Subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate", s.rate)
	fmt.Println("Active?", s.active)
}

func DefaultSubscriber(name string) *Subscriber {
	var s Subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func ApplyDiscount(s *Subscriber) {
	s.rate = 4.99 // 使用点运算符在struct指针和struct上都可访问字段
}
