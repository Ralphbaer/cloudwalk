package app

import (
	"context"
	"log"
)

// App is the application glue where we put all top level components to be used
type App struct {
	*Parser
}

// Run starts the application
// This is the only necessary code to run an app in main.go
func (app *App) Run() {
	log.Printf("QuakeLogParser started")
	if err := app.QuakeLogParserUseCase.Run(context.TODO()); err != nil {
		log.Fatalf("QuakeLogParser exception %#v", err)
	}
}
