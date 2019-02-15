package main

import (
	log "log"
	sync "sync"
	time "time"

	tge "github.com/thommil/tge"
)

type LifeCycle struct {
	Runtime     tge.Runtime
	Counter     int
	totalTick   time.Duration
	totalRender time.Duration
}

func (app *LifeCycle) OnCreate(settings *tge.Settings) error {
	log.Println("OnCreate()")
	settings.Name = "LifeCycle"
	settings.Fullscreen = false
	settings.FPS = 10
	settings.TPS = 10
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
	app.totalRender += elapsedTime
	log.Printf("OnRender(%v) - counter : %d - Total : %d \n", elapsedTime, app.Counter, app.totalRender)
	app.Counter++
	time.Sleep(1 * time.Millisecond)
	locker.Unlock()

	// Simulate heavy treatment
	time.Sleep(4 * time.Millisecond)

}

func (app *LifeCycle) OnTick(elapsedTime time.Duration, locker sync.Locker) {
	// Simulate heavy treatment
	time.Sleep(4 * time.Millisecond)

	// Simulate critical path
	locker.Lock()
	app.totalTick += elapsedTime
	log.Printf("OnTick(%v) - counter : %d - Total : %d \n", elapsedTime, app.Counter, app.totalTick)
	app.Counter++
	time.Sleep(1 * time.Millisecond)
	locker.Unlock()
	// Test stop
	if app.Counter > 100 {
		app.Runtime.Stop()
	}
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
