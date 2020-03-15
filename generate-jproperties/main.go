package main

import (
	"fmt"
	"strings"

	"github.com/keepclean/go-snippets-sandbox/generate-jproperties/backends"
	"github.com/keepclean/go-snippets-sandbox/generate-jproperties/services"
)

func main() {
	serviceSet := []services.Service{
		services.ServiceOne,
		services.ServiceTwo,
		services.ServiceThree,
	}

	printComments()
	allServices(&serviceSet)

	allBackends := make(map[string]backends.Backend)
	for _, s := range serviceSet {
		printService(s, allBackends)
		fmt.Println()
	}

	printBackends(allBackends)
}

func printComments() {
	fmt.Println("#### COMMENTS")
	fmt.Println()
}

func allServices(services *[]services.Service) {
	fmt.Println("#### SERVICES")

	names := make([]string, 0, len(*services))
	for _, s := range *services {
		names = append(names, s.Name)
	}

	fmt.Printf("allServices = %s\n\n", strings.Join(names, ","))
}

func printService(s services.Service, allBackends map[string]backends.Backend) {
	_service := make([]string, 0, len(s.Attempts))

	_attemptsLine := make([]string, 0, len(s.Attempts))
	for _, attempt := range s.Attempts {
		_attemptsLine = append(_attemptsLine, attempt.Number)
	}

	_service = append(
		_service, fmt.Sprintf(
			"service.%s.attempts=%s",
			s.Name,
			strings.Join(_attemptsLine, ","),
		))

	for opt, value := range s.Options {
		_service = append(_service, fmt.Sprintf("service.%s.%s = %v", s.Name, opt, value))
	}

	for _, attempt := range s.Attempts {
		backendInAttempt := make([]string, 0, len(attempt.Backends))

		for _, b := range attempt.Backends {
			if _, ok := allBackends[b.Name]; !ok {
				allBackends[b.Name] = b
			}

			backendInAttempt = append(backendInAttempt, b.Name)
		}

		_service = append(
			_service,
			fmt.Sprintf(
				"service.%s.%s=%s|%s",
				s.Name, attempt.Number,
				strings.Join(backendInAttempt, ","),
				strings.Join(attempt.Sifters, ","),
			),
		)
	}

	fmt.Println(strings.Join(_service, "\n"))
}

func printBackends(allBackends map[string]backends.Backend) {
	fmt.Printf("\n%s\n", "#### BACKENDS")
	for _, b := range allBackends {
		fmt.Printf("backend.%s.path = %s\n", b.Name, b.Path)
		fmt.Printf("backend.%s.connect_timeout = %s\n", b.Name, b.ConnectTimeout)
		fmt.Printf("backend.%s.backend_timeout = %s\n", b.Name, b.BackendTimeout)

		fmt.Println()
	}
}
