package services

import "github.com/keepclean/go-snippets-sandbox/generate-jproperties/backends"

var ServiceTwo = Service{
	Name: "dos",
	Options: map[string]bool{
		"option_1": false,
		"option_2": true,
	},
	Attempts: []Attempt{
		Attempt{
			Number: "0",
			Backends: []backends.Backend{
				backends.BackendOne,
				backends.BackendThree,
			},
			Sifters: []string{"sifter3"},
		},
		Attempt{
			Number: "1",
			Backends: []backends.Backend{
				backends.BackendTwo,
				backends.BackendFour,
			},
			Sifters: []string{"sifter4"},
		},
	},
}
