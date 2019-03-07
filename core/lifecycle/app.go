package main

import (
	fmt "fmt"
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

func (app *LifeCycleApp) OnRender(elapsedTime time.Duration, syncChan <-chan interface{}) {
	<-syncChan
}

func (app *LifeCycleApp) OnTick(elapsedTime time.Duration, syncChan chan<- interface{}) {
	syncChan <- true
}

func (app *LifeCycleApp) OnPause() {
	fmt.Println("OnPause()")
}

func (app *LifeCycleApp) OnStop() {
	fmt.Println("OnStop()")
}

func (app *LifeCycleApp) OnDispose() {
	fmt.Println("OnDispose()")
}

func main() {
	tge.Run(&LifeCycleApp{})
}
