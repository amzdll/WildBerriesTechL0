package validator

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"validator",
		fx.Provide(validator.New),
	)
}
