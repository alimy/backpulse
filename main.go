package main

import (
	"github.com/backpulse/core/cmd/core"

	_ "github.com/backpulse/core/cmd/serve"
	_ "github.com/backpulse/core/cmd/version"
)

func main() {
	// Setup root cli command of application
	core.Setup(
		"backpulse",                 // command name
		"Provide backpulse service", // command short describe
		"Provide backpulse service", // command long describe
	)

	// Execute start application
	core.Execute()
}
