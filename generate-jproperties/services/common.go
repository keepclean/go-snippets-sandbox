package services

import "github.com/keepclean/go-snippets-sandbox/generate-jproperties/backends"

type Service struct {
	Name     string
	Options  map[string]bool
	Attempts []Attempt
}

type Attempt struct {
	Number   string
	Backends []backends.Backend
	Sifters  []string
}
