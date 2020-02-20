package main

import "fmt"

type Message interface {
	Message() string
	Destination() string
	Sender() string
}

type SMS struct {
	From string
	To   string
	Msg  string
}

func NewSMS(origin string, dest string, message string) *SMS {
	return &SMS{
		From: origin,
		To:   dest,
		Msg:  message,
	}
}

func (s *SMS) Message() string {
	return s.Msg
}

func (s *SMS) Sender() string {
	return s.From
}

func (s *SMS) Destination() string {
	return s.To
}

type Address struct {
	StreetName string
	Zipcode    string
}

func (a *Address) String() string {
	return fmt.Sprintf("%v -- %v", a.StreetName, a.Zipcode)
}

func NewAddress(streetName, zipcode string) *Address {
	return &Address{
		StreetName: streetName,
		Zipcode:    zipcode,
	}
}

type Pigeon struct {
	From string
	Msg  string
	To   *Address
}

func NewPigeon(sender string, message string, addr *Address) *Pigeon {
	return &Pigeon{
		From: sender,
		Msg:  message,
		To:   addr,
	}
}

func (p *Pigeon) Message() string {
	return p.Msg
}

func (p *Pigeon) Sender() string {
	return p.From
}

func (p *Pigeon) Destination() string {
	return fmt.Sprintf("%v", p.To)
}
