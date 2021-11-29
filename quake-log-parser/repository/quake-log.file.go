package repository

import (
	"context"
	"os"
)

// QuakeLogFileRepository represents a implementation of QuakeLogRepository interface
type QuakeLogFileRepository struct {
	FilePath string
}

// NewQuakeLogFileRepository creates an instance of QuakeLogFileRepository
func NewQuakeLogFileRepository(filePath string) *QuakeLogFileRepository {
	return &QuakeLogFileRepository{
		FilePath: filePath,
	}
}

// GetFile returns the file object
func (c *QuakeLogFileRepository) GetFile(ctx context.Context) (*os.File, error) {
	f, err := os.Open(c.FilePath)
	if err != nil {
		return nil, err
	}

	return f, nil
}
