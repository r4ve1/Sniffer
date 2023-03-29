package reader

import (
	"sniffer/internal/cache"
)

type I interface {
	Start(filter string) (cache.I, error)
	Stop() error
}
