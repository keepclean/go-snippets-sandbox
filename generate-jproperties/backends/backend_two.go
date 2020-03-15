package backends

import "time"

var BackendTwo = Backend{
	Name:           "backend2",
	Path:           "path2",
	ConnectTimeout: 100 * time.Millisecond,
	BackendTimeout: 300 * time.Millisecond,
}
