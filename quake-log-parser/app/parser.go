package app

import (
	uc "github.com/Ralphbaer/cloudwalk/quake-log-parser/usecase"
)

// Parser represents the daemon processs for member-glucose service
type Parser struct {
	QuakeLogParserUseCase *uc.QuakeLogParserUseCase
}

// NewDaemon creates an instance of Daemon
func NewParser(usecase *uc.QuakeLogParserUseCase) *Parser {
	return &Parser{
		QuakeLogParserUseCase: usecase,
	}
}
