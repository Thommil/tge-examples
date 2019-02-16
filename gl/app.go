package main

import (
	log "log"
	sync "sync"
	time "time"

	tge "github.com/thommil/tge"
	gl "github.com/thommil/tge-gl"
)

type LifeCycle struct {
	Runtime tge.Runtime
}

func (app *LifeCycle) OnCreate(settings *tge.Settings) error {
	log.Println("OnCreate()")
	settings.Name = "GL"
	settings.Fullscreen = false
	settings.FPS = 10
	settings.TPS = 10
	return nil
}

func (app *LifeCycle) OnStart(runtime tge.Runtime) error {
	log.Println("OnStart()")
	app.Runtime = runtime

	err := gl.Init(runtime)
	if err != nil {
		return err
	}

	return nil
}

func (app *LifeCycle) OnResize(width int, height int) {
	log.Printf("OnResize(%d, %d)\n", width, height)
}

func (app *LifeCycle) OnResume() {
	log.Println("OnResume()")
	gl.ClearColor(0.15, 0.04, 0.15, 1)
}

func (app *LifeCycle) OnRender(elapsedTime time.Duration, locker sync.Locker) {
	log.Printf("OnRender(%v)\n", elapsedTime)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (app *LifeCycle) OnTick(elapsedTime time.Duration, locker sync.Locker) {
	log.Printf("OnTick(%v)\n", elapsedTime)

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
