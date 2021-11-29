package repository

import (
	"context"
	"os"
)

// QuakeLogJSONRepository represents a local JSON implementation of QuakeLogRepository interface
type QuakeLogFileRepository struct {
	FilePath string
}

// NewQuakeLogFileRepository creates an instance of repository.QuakeLogJSONRepository
func NewQuakeLogFileRepository(filePath string) *QuakeLogFileRepository {
	return &QuakeLogFileRepository{
		FilePath: filePath,
	}
}

// List returns a list of QuakeLogs given the QuakeLogFilter type parameter
func (c *QuakeLogFileRepository) GetFile(ctx context.Context) (*os.File, error) {
	f, err := os.Open(c.FilePath)
	if err != nil {
		return nil, err
	}

	return f, nil
}
