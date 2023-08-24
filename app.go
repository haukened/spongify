package main

import (
	"context"
	"fmt"
	"log"
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
	height, width, err := getPrimaryScreenDimensions(ctx)
	if err != nil {
		log.Print(err)
	}
	// calculate the position of the window
	if height != 0 && width != 0 {
		X_POS := width - WINDOW_WIDTH - 50
		Y_POS := height - WINDOW_HEIGHT - 150
		// and position the window
		wuntime.WindowSetPosition(ctx, X_POS, Y_POS)
	}
	// emit the navigation event to load the about page in the frontend
	wuntime.EventsEmit(ctx, "navigate", "/about")
	// give the page time to load
	wuntime.WindowShow(ctx)
}

func getPrimaryScreenDimensions(c context.Context) (height int, width int, err error) {
	screens, err := wuntime.ScreenGetAll(c)
	if err != nil {
		return
	}
	for _, screen := range screens {
		if screen.IsPrimary {
			height = screen.Height
			width = screen.Width
			return
		}
	}
	err = fmt.Errorf("unable to detect primary screen")
	return
}
