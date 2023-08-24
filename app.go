package main

import (
	"context"
	"fmt"
	"runtime"
	"spongify/internal/config"

	"fyne.io/systray"
	wuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx      context.Context
	conf     *config.Config
	trayExit func()
}

// NewApp creates a new App application struct
func NewApp() *App {
	c, err := config.Load(APP_NAME)
	if err != nil {
		panic(err)
	}
	return &App{
		conf: c,
	}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	// set up the systray
	start, stop := systray.RunWithExternalLoop(setupSystray(a.ctx), nil)
	a.trayExit = stop
	// run the systray
	start()
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	if a.trayExit != nil {
		a.trayExit()
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func setupSystray(ctx context.Context) func() {
	return func() {
		// set the systray icons and settings
		if runtime.GOOS == "windows" {
			systray.SetTemplateIcon(iconWin, iconWin)
		} else {
			systray.SetTemplateIcon(icon, icon)
		}
		systray.SetTitle(APP_DISPLAY_NAME)
		systray.SetTooltip(APP_DISPLAY_NAME)

		// add menu items
		mAbout := systray.AddMenuItem("About", "About this app")
		mAbout.Enable()
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Quit", "Exit the app")
		mQuit.Enable()

		// background task to listen for click events
		go func() {
			for {
				select {
				case <-mQuit.ClickedCh:
					// exit the systray
					systray.Quit()
					// exit the app
					wuntime.Quit(ctx)
				case <-mAbout.ClickedCh:
					showAbout(ctx)
				}
			}
		}()
	}
}

func showAbout(ctx context.Context) {
	// here we need to navigate to the about page and then show it
}
