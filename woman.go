package main

import (
	"errors"
	"fmt"
)

type WomanConfig struct {
	Name                 string
	PhoneNumber          string
	Address              *Address
	BoyfriendName        string
	BoyfriendPhoneNumber string
}

type Woman struct {
	Cfg WomanConfig
}

// Provider
func NewWoman(cfg WomanConfig) *Woman {
	if cfg.Name == "" {
		cfg.Name = "Woman"
	}

	return &Woman{
		Cfg: cfg,
	}
}

func (w *Woman) ReadMessage(msg Message) (err error) {
	if w.Cfg.BoyfriendName != msg.Sender() && w.Cfg.BoyfriendPhoneNumber != msg.Sender() {
		err = errors.New("woman: This is not my boyfriend")
		fmt.Printf("%s: Who is this? This is not my boyfriend\n", w.Cfg.Name)
		return
	}

	fmt.Println(fmt.Sprintf("%s: This is from my dear, %s", w.Cfg.Name, msg.Sender()))
	fmt.Println(fmt.Sprintf("%s: The message is: %s", w.Cfg.Name, msg.Message()))
	return
}

func (w *Woman) SendMessage(msg Message) {
	fmt.Println("I dont know how to send message...")
}

func (w *Woman) Name() string {
	return w.Cfg.Name
}

func (w *Woman) Address() *Address {
	return w.Cfg.Address
}

func (w *Woman) PhoneNumber() string {
	return w.Cfg.PhoneNumber
}
