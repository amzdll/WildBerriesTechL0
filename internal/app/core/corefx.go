package core

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"Core Module",
		orderModule(),
	)
}
