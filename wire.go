//+build wireinject

package main

import (
	"context"

	"github.com/google/wire"
)

// var otpRepoBinding = wire.Bind(new(repository.OTPRepoer), new(*repository.OTPRepo))
// var otpValidatorBinding = wire.Bind(new(service.OTPValidator), new(*service.Validator))

var manWoman = wire.NewSet(NewMan, NewWoman)

func InitializeMan(ctx context.Context, manName string, wcfg WomanConfig) *Man {
	// panic(wire.Build(NewMan, NewWoman))
	panic(wire.Build(manWoman))
}
