package main

import (
	fmt "fmt"
	sync "sync"
	time "time"

	tge "github.com/thommil/tge"
)

type LifeCycleApp struct {
	Runtime     tge.Runtime
	Counter     int
	totalTick   time.Duration
	totalRender time.Duration
}

func (app *LifeCycleApp) OnCreate(settings *tge.Settings) error {
	fmt.Println("OnCreate()")
	settings.Name = "LifeCycleApp"
	settings.Fullscreen = false
	settings.TPS = 10
	settings.EventMask = tge.AllEventsDisable
	return nil
}

func (app *LifeCycleApp) OnStart(runtime tge.Runtime) error {
	fmt.Println("OnStart()")
	app.Runtime = runtime
	return nil
}

func (app *LifeCycleApp) OnResize(width int, height int) {
	fmt.Printf("OnResize(%d, %d)\n", width, height)
}

func (app *LifeCycleApp) OnResume() {
	fmt.Println("OnResume()")
}

func (app *LifeCycleApp) OnRender(elapsedTime time.Duration, mutex *sync.Mutex) {
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

func (app *LifeCycleApp) OnTick(elapsedTime time.Duration, mutex *sync.Mutex) {
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

func (app *LifeCycleApp) OnMouseEvent(event tge.MouseEvent) {
	// NOP
}

func (app *LifeCycleApp) OnScrollEvent(event tge.ScrollEvent) {
	// NOP
}

func (app *LifeCycleApp) OnKeyEvent(event tge.KeyEvent) {
	// NOP
}

func (app *LifeCycleApp) OnPause() {
	fmt.Println("OnPause()")
}

func (app *LifeCycleApp) OnStop() {
	fmt.Println("OnStop()")
}

func (app *LifeCycleApp) OnDispose() error {
	fmt.Println("OnDispose()")
	return nil
}

func main() {
	tge.Run(&LifeCycleApp{})
}
