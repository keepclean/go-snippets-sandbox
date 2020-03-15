package backends

import "time"

var BackendFour = Backend{
	Name:           "backend4",
	Path:           "path4",
	ConnectTimeout: 100 * time.Millisecond,
	BackendTimeout: 300 * time.Millisecond,
}
