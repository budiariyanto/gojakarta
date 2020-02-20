package main

import "fmt"

type Human interface {
	SendMessage(msg Message)
	ReadMessage(msg Message) error
	Name() string
	Address() *Address
	PhoneNumber() string
}

type Man struct {
	Name        string
	PhoneNumber string
	Woman       Human
}

func NewMan(name string, w Human) *Man {
	if name == "" {
		name = "Man"
	}

	return &Man{
		Name:  name,
		Woman: w,
	}
}

func (m *Man) SendMessage(msg Message) {
	fmt.Println(fmt.Sprintf("%s: I'm sending message to my girlfriend %s at %v", m.Name, m.Woman.Name(), msg.Destination()))

	if msg.Destination() != fmt.Sprintf("%v", m.Woman.Address()) {
		fmt.Println("Message not sent! So sad... :(")
		return
	}

	m.Woman.ReadMessage(msg)
}
