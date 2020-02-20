//+build wireinject

package main

import (
	"context"

	"github.com/google/wire"
)

func InitializeMan(ctx context.Context, manName string, wcfg WomanConfig) *Man {
	panic(wire.Build(NewMan, NewWoman))
}
