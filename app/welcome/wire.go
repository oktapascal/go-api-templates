//go:build wireinject
// +build wireinject

package welcome

import (
	"github.com/google/wire"
)

func Wire() *Handler {
	panic(wire.Build(ProviderSet))
}
