package backends

import "time"

var BackendThree = Backend{
	Name:           "backend3",
	Path:           "path3",
	ConnectTimeout: 100 * time.Millisecond,
	BackendTimeout: 300 * time.Millisecond,
}
