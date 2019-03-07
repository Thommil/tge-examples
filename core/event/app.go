package main

import (
	fmt "fmt"
	time "time"

	tge "github.com/thommil/tge"
)

type EventApp struct {
	runtime tge.Runtime
}

type TestEvent struct{}

func (t TestEvent) Channel() string {
	return "propagate-test"
}

func (app *EventApp) OnCreate(settings *tge.Settings) error {
	fmt.Println("OnCreate()")
	settings.Name = "EventApp"
	settings.Fullscreen = false
	settings.EventMask = tge.AllEventsEnabled
	return nil
}

func (app *EventApp) OnStart(runtime tge.Runtime) error {
	fmt.Println("OnStart()")
	app.runtime = runtime
	runtime.Subscribe(tge.ResizeEvent{}.Channel(), app.OnResize)
	runtime.Subscribe(tge.MouseEvent{}.Channel(), app.OnMouseEvent)
	runtime.Subscribe(tge.ScrollEvent{}.Channel(), app.OnScrollEvent)
	runtime.Subscribe(tge.KeyEvent{}.Channel(), app.OnKeyEvent)

	runtime.Subscribe(TestEvent{}.Channel(), func(event tge.Event) bool {
		fmt.Println("First listener -> propagate")
		return false
	})

	runtime.Subscribe(TestEvent{}.Channel(), func(event tge.Event) bool {
		fmt.Println("Second listener -> stop propagate")
		return true
	})

	runtime.Subscribe(TestEvent{}.Channel(), func(event tge.Event) bool {
		fmt.Println("Third listener -> SHOULD NEVER BE DISPLAYED ERROR ERROR")
		return false
	})

	return nil
}

func (app *EventApp) OnResume() {
	fmt.Println("OnResume()")
}

func (app *EventApp) OnRender(elapsedTime time.Duration, syncChan <-chan interface{}) {
	<-syncChan
}

func (app *EventApp) OnTick(elapsedTime time.Duration, syncChan chan<- interface{}) {
	syncChan <- true
}

func (app *EventApp) OnResize(event tge.Event) bool {
	fmt.Printf("OnResize() : %v\n", event)
	return false
}

func (app *EventApp) OnMouseEvent(event tge.Event) bool {
	fmt.Printf("OnMouseEvent() : %v\n", event)
	return false
}

func (app *EventApp) OnScrollEvent(event tge.Event) bool {
	fmt.Printf("OnScrollEvent() : %v\n", event)
	return false
}

func (app *EventApp) OnKeyEvent(event tge.Event) bool {
	e := event.(tge.KeyEvent)
	if e.Type == tge.TypeDown {
		fmt.Println("OnKeyEvent :")
		fmt.Printf("   Value : %s\n", e.Value)
		fmt.Printf("   IsValid : %v\n", e.Key.IsValid())
		fmt.Printf("   IsPrintable : %v\n", e.Key.IsPrintable())
		fmt.Printf("   IsAction : %v\n", e.Key.IsAction())
		fmt.Printf("   IsFunction : %v\n", e.Key.IsFunction())
		fmt.Printf("   IsModifier : %v\n", e.Key.IsModifier())
		fmt.Printf("   IsCompose : %v\n", e.Key.IsCompose())
	} else {
		if e.Key == tge.KeyCodeEscape {
			app.runtime.Stop()
		}
		if e.Key == tge.KeyCodeSpacebar {
			app.runtime.Publish(TestEvent{})
		}
	}

	return false
}

func (app *EventApp) OnPause() {
	fmt.Println("OnPause()")
}

func (app *EventApp) OnStop() {
	fmt.Println("OnStop()")
	app.runtime.Subscribe(tge.ResizeEvent{}.Channel(), app.OnResize)
	app.runtime.Unsubscribe(tge.MouseEvent{}.Channel(), app.OnMouseEvent)
	app.runtime.Unsubscribe(tge.ScrollEvent{}.Channel(), app.OnScrollEvent)
	app.runtime.Unsubscribe(tge.KeyEvent{}.Channel(), app.OnKeyEvent)
}

func (app *EventApp) OnDispose() {
	fmt.Println("OnDispose()")
}

func main() {
	tge.Run(&EventApp{})
}
