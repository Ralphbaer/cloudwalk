package common

import (
	"sync"
)

// App represents an application that will run as a deployable component.
// It's an entrypoint at main.go
type App interface {
	Run(launcher *Launcher) error
}

//LauncherOption defines a function option for Launcher
type LauncherOption func(l *Launcher)

// RunApp start all process registered before to the launcher
func RunApp(name string, app App) LauncherOption {
	return func(l *Launcher) {
		l.Add(name, app)
	}
}

// Launcher manages apps
type Launcher struct {
	apps    map[string]App
	wg      *sync.WaitGroup
	Verbose bool
}

// Add runs an applicaton in a goroutine.
func (l *Launcher) Add(appName string, a App) *Launcher {
	l.apps[appName] = a
	return l
}

// Run run every application registered before with Run method.
func (l *Launcher) Run() {
	count := len(l.apps)
	l.wg.Add(count)

	for name, app := range l.apps {
		go func(name string, app App) {
			l.wg.Done()
		}(name, app)
	}

	l.wg.Wait()
}

// NewLauncher create an instance of Launch
func NewLauncher(opts ...LauncherOption) *Launcher {

	l := &Launcher{
		apps:    make(map[string]App),
		wg:      new(sync.WaitGroup),
		Verbose: true,
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}
