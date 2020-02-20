//+build wireinject

package main

import (
	"context"

	"github.com/google/wire"
)

var womanBinding = wire.Bind(new(Human), new(*Woman))

var manWoman = wire.NewSet(NewMan, NewWoman, womanBinding)

func InitializeMan(ctx context.Context, manName string, wcfg WomanConfig) *Man {
	panic(wire.Build(manWoman))
}
