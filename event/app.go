package main

import (
	fmt "fmt"
	sync "sync"
	time "time"

	tge "github.com/thommil/tge"
)

type EventApp struct {
	Runtime tge.Runtime
}

func (app *EventApp) OnCreate(settings *tge.Settings) error {
	fmt.Println("OnCreate()")
	settings.Name = "EventApp"
	settings.Fullscreen = false
	settings.FPS = 100
	settings.TPS = 1
	settings.EventMask = tge.MouseEventEnabled | tge.ScrollEventEnabled | tge.KeyEventEnabled
	return nil
}

func (app *EventApp) OnStart(runtime tge.Runtime) error {
	fmt.Println("OnStart()")
	app.Runtime = runtime
	return nil
}

func (app *EventApp) OnResize(width int, height int) {
	fmt.Printf("OnResize(%d, %d)\n", width, height)
}

func (app *EventApp) OnResume() {
	fmt.Println("OnResume()")
}

func (app *EventApp) OnRender(elapsedTime time.Duration, mutex *sync.Mutex) {
}

func (app *EventApp) OnTick(elapsedTime time.Duration, mutex *sync.Mutex) {
}

func (app *EventApp) OnMouseEvent(event tge.MouseEvent) {
	fmt.Printf("OnMouseEvent() : %v\n", event)
}

func (app *EventApp) OnScrollEvent(event tge.ScrollEvent) {
	fmt.Printf("OnScrollEvent() : %v\n", event)
}

func (app *EventApp) OnKeyEvent(event tge.KeyEvent) {
	fmt.Printf("OnKeyEvent() : %v\n", event)
}

func (app *EventApp) OnPause() {
	fmt.Println("OnPause()")
}

func (app *EventApp) OnStop() {
	fmt.Println("OnStop()")
}

func (app *EventApp) OnDispose() error {
	fmt.Println("OnDispose()")
	return nil
}

func main() {
	tge.Run(&EventApp{})
}
