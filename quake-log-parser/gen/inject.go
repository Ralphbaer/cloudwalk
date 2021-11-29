//go:build wireinject
// +build wireinject

//golint:ignore

package gen

import (
	"fmt"
	"os"

	"github.com/Ralphbaer/cloudwalk/quake-log-parser/app"
	"github.com/google/wire"

	r "github.com/Ralphbaer/cloudwalk/quake-log-parser/repository"
	uc "github.com/Ralphbaer/cloudwalk/quake-log-parser/usecase"
)

func setupFileLocation() string {
	pwd, _ := os.Getwd()
	return fmt.Sprintf("%s/gen/qgames.log", pwd)
}

var applicationSet = wire.NewSet(
	app.NewParser,
	setupFileLocation,
	r.NewQuakeLogFileRepository,
	wire.Struct(new(uc.QuakeLogParserUseCase), "*"),
	wire.Bind(new(r.QuakeLogRepository), new(*r.QuakeLogFileRepository)),
)

// InitializeApp setup the dependencies and returns a new *app.App instance
func InitializeApp() *app.App {
	wire.Build(
		applicationSet,
		wire.Struct(new(app.App), "*"),
	)
	return nil
}

// InitializeMeasurementUseCase setup the dependencies and returns a new *usecase.ReportUseCase
func InitializeUseCase() *uc.QuakeLogParserUseCase {
	wire.Build(
		applicationSet,
	)
	return nil
}
