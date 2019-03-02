package main

import (
	fmt "fmt"
	sync "sync"
	time "time"

	tge "github.com/thommil/tge"
)

type LifeCycleApp struct {
	Runtime tge.Runtime
}

func (app *LifeCycleApp) OnCreate(settings *tge.Settings) error {
	fmt.Println("OnCreate()")
	settings.Name = "LifeCycleApp"
	settings.Fullscreen = false
	settings.TPS = 1
	settings.EventMask = tge.AllEventsDisable
	return nil
}

func (app *LifeCycleApp) OnStart(runtime tge.Runtime) error {
	fmt.Println("OnStart()")
	app.Runtime = runtime
	return nil
}

func (app *LifeCycleApp) OnResume() {
	fmt.Println("OnResume()")
}

func (app *LifeCycleApp) OnRender(elapsedTime time.Duration, mutex *sync.Mutex) {
	mutex.Lock()
	fmt.Println("OnRender()")
	mutex.Unlock()
}

func (app *LifeCycleApp) OnTick(elapsedTime time.Duration, mutex *sync.Mutex) {
	mutex.Lock()
	fmt.Println("OnTick()")
	mutex.Unlock()
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
