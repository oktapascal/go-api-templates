package welcome

import (
	"github.com/google/wire"
	"sync"
)

var (
	hdl     *Handler
	hdlOnce sync.Once

	ProviderSet = wire.NewSet(ProvideHandler)
)

func ProvideHandler() *Handler {
	hdlOnce.Do(func() {
		hdl = &Handler{}
	})

	return hdl
}
