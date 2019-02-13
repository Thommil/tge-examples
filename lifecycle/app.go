package main

import (
	log "log"
	sync "sync"
	time "time"

	tge "github.com/thommil/tge"
)

type LifeCycle struct {
	Runtime tge.Runtime
	Counter int
}

func (app *LifeCycle) OnCreate(settings *tge.Settings) error {
	log.Println("OnCreate()")
	settings.Name = "LifeCycle"
	settings.Fullscreen = false
	settings.FPS = 1
	settings.TPS = 2
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

func (app *LifeCycle) OnRender(elapsedTime time.Duration, locker sync.Locker) {
	// Simulate critical path
	locker.Lock()
	log.Printf("OnRender(%v)\n", elapsedTime)
	time.Sleep(50 * time.Millisecond)
	app.Counter++
	log.Printf("Render -> Counter : (%v)\n", app.Counter)
	locker.Unlock()

	// Simulate heavy treatment
	time.Sleep(500 * time.Millisecond)

	// Test stop
	if app.Counter > 20 {
		app.Runtime.Stop()
	}

}

func (app *LifeCycle) OnTick(elapsedTime time.Duration, locker sync.Locker) {
	// Simulate heavy treatment
	time.Sleep(100 * time.Millisecond)

	// Simulate critical path
	locker.Lock()
	log.Printf("OnTick(%v)\n", elapsedTime)
	time.Sleep(50 * time.Millisecond)
	app.Counter++
	log.Printf("Ticker -> Counter : (%v)\n", app.Counter)
	locker.Unlock()
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
