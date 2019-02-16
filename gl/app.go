package main

import (
	fmt "fmt"
	sync "sync"
	time "time"

	tge "github.com/thommil/tge"
	gl "github.com/thommil/tge-gl"
)

type LifeCycle struct {
	Runtime tge.Runtime
}

func (app *LifeCycle) OnCreate(settings *tge.Settings) error {
	fmt.Println("OnCreate()")
	settings.Name = "GL"
	settings.Fullscreen = false
	settings.FPS = 10
	settings.TPS = 10

	return nil
}

func (app *LifeCycle) OnStart(runtime tge.Runtime) error {
	fmt.Println("OnStart()")
	app.Runtime = runtime

	runtime.Use(gl.GetPlugin())

	return nil
}

func (app *LifeCycle) OnResize(width int, height int) {
	fmt.Printf("OnResize(%d, %d)\n", width, height)
}

func (app *LifeCycle) OnResume() {
	fmt.Println("OnResume()")
	gl.ClearColor(0.15, 0.04, 0.15, 1)
}

func (app *LifeCycle) OnRender(elapsedTime time.Duration, locker sync.Locker) {
	fmt.Printf("OnRender(%v)\n", elapsedTime)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (app *LifeCycle) OnTick(elapsedTime time.Duration, locker sync.Locker) {
	fmt.Printf("OnTick(%v)\n", elapsedTime)

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
