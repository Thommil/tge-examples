package main

import (
	fmt "fmt"
	time "time"

	tge "github.com/thommil/tge"
	gesture "github.com/thommil/tge-gesture"
)

type GestureApp struct {
	runtime tge.Runtime
}

func (app *GestureApp) OnCreate(settings *tge.Settings) error {
	fmt.Println("OnCreate()")
	settings.Name = "GestureApp"
	settings.Fullscreen = false
	settings.EventMask = tge.AllEventsEnabled
	return nil
}

func (app *GestureApp) OnStart(runtime tge.Runtime) error {
	fmt.Println("OnStart()")
	app.runtime = runtime
	runtime.Subscribe(gesture.PinchEvent{}.Channel(), app.OnPinch)
	runtime.Subscribe(gesture.SwipeEvent{}.Channel(), app.OnSwipe)
	runtime.Subscribe(gesture.LongPressEvent{}.Channel(), app.OnLongPess)
	return nil
}

func (app *GestureApp) OnResume() {
	fmt.Println("OnResume()")
}

func (app *GestureApp) OnPinch(event tge.Event) bool {
	fmt.Printf("OnPinch() : %v\n", event)
	return false
}

func (app *GestureApp) OnSwipe(event tge.Event) bool {
	fmt.Printf("OnSwipe() : %v\n", event)
	return false
}

func (app *GestureApp) OnLongPess(event tge.Event) bool {
	fmt.Printf("OnLongPess() : %v\n", event)
	return false
}

func (app *GestureApp) OnRender(elapsedTime time.Duration, syncChan <-chan interface{}) {
	<-syncChan
}

func (app *GestureApp) OnTick(elapsedTime time.Duration, syncChan chan<- interface{}) {
	syncChan <- true
}

func (app *GestureApp) OnPause() {
	fmt.Println("OnPause()")
}

func (app *GestureApp) OnStop() {
	fmt.Println("OnStop()")
	app.runtime.Unsubscribe(gesture.PinchEvent{}.Channel(), app.OnPinch)
	app.runtime.Unsubscribe(gesture.SwipeEvent{}.Channel(), app.OnSwipe)
	app.runtime.Unsubscribe(gesture.LongPressEvent{}.Channel(), app.OnLongPess)
}

func (app *GestureApp) OnDispose() {
	fmt.Println("OnDispose()")
}

func main() {
	tge.Run(&GestureApp{})
}
