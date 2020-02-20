package main

import (
	"context"
)

type MessageType string

const (
	SMSType    MessageType = "sms"
	PigeonType MessageType = "pigeon"
)

func main() {
	ctx := context.Background()
	wcfg := WomanConfig{
		Name:                 "Siti",
		BoyfriendName:        "Tono",
		BoyfriendPhoneNumber: "08123456789",
		Address:              NewAddress("Jl. Kebenaran no. 101", "88888"),
	}

	man := InitializeMan(ctx, "Toni", wcfg)

	messageType := PigeonType
	var msg Message

	switch messageType {
	case SMSType:
		msg = NewSMS(man.PhoneNumber, man.Woman.PhoneNumber(), "Hi dear...")
	case PigeonType:
		msg = NewPigeon(man.Name, "Hi dear...", man.Woman.Address())
	}

	man.SendMessage(msg)
}
