package backends

import "time"

var BackendOne = Backend{
	Name:           "backend1",
	Path:           "path1",
	ConnectTimeout: 100 * time.Millisecond,
	BackendTimeout: 300 * time.Millisecond,
}
