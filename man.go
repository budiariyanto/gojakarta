package main

import "fmt"

type Man struct {
	Name        string
	PhoneNumber string
	Woman       *Woman
}

func NewMan(name string, w *Woman) *Man {
	if name == "" {
		name = "Man"
	}

	return &Man{
		Name:  name,
		Woman: w,
	}
}

func (m *Man) SendMessage(msg Message) {
	fmt.Println(fmt.Sprintf("%s: I'm sending message to my girlfriend %s at %v", m.Name, m.Woman.Cfg.Name, msg.Destination()))

	if msg.Destination() != fmt.Sprintf("%v", m.Woman.Cfg.Address) {
		fmt.Println("Message not sent! So sad... :(")
		return
	}

	m.Woman.ReadMessage(msg)
}
