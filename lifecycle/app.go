package main

import (
	log "log"
	"time"

	tge "github.com/thommil/tge"
)

type LifeCycle struct {
	Runtime    tge.Runtime
	FPSCounter int
	TPSCounter int
}

func (app *LifeCycle) OnCreate(settings *tge.Settings) error {
	log.Println("OnCreate()")
	settings.Name = "LifeCycle"
	settings.Fullscreen = false
	settings.Renderer.FPS = 2
	settings.Physics.TPS = 4
	return nil
}

func (app *LifeCycle) OnStart(runtime tge.Runtime) error {
	log.Println("OnStart()")
	app.Runtime = runtime
	return nil
}

func (app *LifeCycle) OnResize(width int, height int) {
	log.Printf("OnResize(%d, %d)\n", width, height)
}

func (app *LifeCycle) OnResume() {
	log.Println("OnResume()")
}

func (app *LifeCycle) OnRender(elapsedTime time.Duration) {
	log.Printf("OnRender(%v)\n", elapsedTime)
	app.FPSCounter++
	time.Sleep(100 * time.Millisecond)
	if app.FPSCounter > 10 {
		app.Runtime.Stop()
	}
}

func (app *LifeCycle) OnTick(elapsedTime time.Duration) {
	log.Printf("OnTick(%v)\n", elapsedTime)
	app.TPSCounter++
}

func (app *LifeCycle) OnPause() {
	log.Println("OnPause()")
}

func (app *LifeCycle) OnStop() {
	log.Println("OnStop()")
}

func (app *LifeCycle) OnDispose() error {
	log.Println("OnDispose()")
	return nil
}

func main() {
	tge.Run(&LifeCycle{})
}
