package services

import "github.com/keepclean/go-snippets-sandbox/generate-jproperties/backends"

var ServiceOne = Service{
	Name: "uno",
	Options: map[string]bool{
		"option_1": true,
		"option_2": false,
	},
	Attempts: []Attempt{
		Attempt{
			Number: "0",
			Backends: []backends.Backend{
				backends.BackendOne,
				backends.BackendTwo,
			},
			Sifters: []string{"sifter1"},
		},
		Attempt{
			Number: "1",
			Backends: []backends.Backend{
				backends.BackendThree,
				backends.BackendFour,
			},
			Sifters: []string{"sifter2"},
		},
	},
}
