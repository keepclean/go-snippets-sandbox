package services

var ServiceThree = Service{
	Name: "tres",
	Options: map[string]bool{
		"option_1": true,
		"option_2": false,
		"option_3": false,
	},
	Attempts: ServiceOne.Attempts,
}
