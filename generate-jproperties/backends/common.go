package backends

import "time"

type Backend struct {
	Name           string
	Path           string
	ConnectTimeout time.Duration
	BackendTimeout time.Duration
}
