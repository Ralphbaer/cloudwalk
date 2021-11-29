package repository

import (
	"context"
	"os"
)

//go:generate mockgen -destination=../gen/mock/repository_mock.go -package=mock . QuakeLogRepository

// QuakeLogRepository manages quake-log-file repository operations
type QuakeLogRepository interface {
	GetFile(ctx context.Context) (*os.File, error)
}
