package app

import (
	uc "github.com/Ralphbaer/cloudwalk/quake-log-parser/usecase"
)

// Parser represents the process for quake-log-parser
type Parser struct {
	QuakeLogParserUseCase *uc.QuakeLogParserUseCase
}

// NewParser creates an instance of Parser
func NewParser(usecase *uc.QuakeLogParserUseCase) *Parser {
	return &Parser{
		QuakeLogParserUseCase: usecase,
	}
}
