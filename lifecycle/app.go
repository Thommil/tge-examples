package main

import (
	fmt "fmt"
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
	fmt.Println("OnCreate()")
	settings.Name = "LifeCycle"
	settings.Fullscreen = false
	settings.FPS = 10
	settings.TPS = 10
	return nil
}

func (app *LifeCycle) OnStart(runtime tge.Runtime) error {
	fmt.Println("OnStart()")
	app.Runtime = runtime
	return nil
}

func (app *LifeCycle) OnResize(width int, height int) {
	fmt.Printf("OnResize(%d, %d)\n", width, height)
}

func (app *LifeCycle) OnResume() {
	fmt.Println("OnResume()")
}

func (app *LifeCycle) OnRender(elapsedTime time.Duration, mutex *sync.Mutex) {
	// Simulate critical path
	mutex.Lock()
	app.totalRender += elapsedTime
	fmt.Printf("OnRender(%v) - counter : %d - Total : %d \n", elapsedTime, app.Counter, app.totalRender)
	app.Counter++
	time.Sleep(1 * time.Millisecond)
	mutex.Unlock()

	// Simulate heavy treatment
	time.Sleep(4 * time.Millisecond)

	// Test stop
	if app.Counter > 100 {
		app.Runtime.Stop()
	}
}

func (app *LifeCycle) OnTick(elapsedTime time.Duration, mutex *sync.Mutex) {
	// Simulate heavy treatment
	time.Sleep(4 * time.Millisecond)

	// Simulate critical path
	mutex.Lock()
	app.totalTick += elapsedTime
	fmt.Printf("OnTick(%v) - counter : %d - Total : %d \n", elapsedTime, app.Counter, app.totalTick)
	app.Counter++
	time.Sleep(1 * time.Millisecond)
	mutex.Unlock()

}

func (app *LifeCycle) OnPause() {
	fmt.Println("OnPause()")
}

func (app *LifeCycle) OnStop() {
	fmt.Println("OnStop()")
}

func (app *LifeCycle) OnDispose() error {
	fmt.Println("OnDispose()")
	return nil
}

func main() {
	tge.Run(&LifeCycle{})
}
