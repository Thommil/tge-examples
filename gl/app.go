package main

import (
	fmt "fmt"
	sync "sync"
	time "time"

	tge "github.com/thommil/tge"
	gl "github.com/thommil/tge-gl"
)

type LifeCycle struct {
	Runtime        tge.Runtime
	TPSStartTime   time.Time
	TPSPrinterTime time.Time
	FPSStartTime   time.Time
	FPSPrinterTime time.Time
	TPSCounter     float64
	FPSCounter     float64
}

func (app *LifeCycle) OnCreate(settings *tge.Settings) error {
	fmt.Println("OnCreate()")
	settings.Name = "GL"
	settings.Fullscreen = true
	settings.FPS = 100
	settings.TPS = 100

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
	app.TPSStartTime = time.Now()
	app.FPSStartTime = time.Now()
	app.FPSPrinterTime = time.Now().Add(10 * time.Second)
	app.TPSPrinterTime = time.Now().Add(10 * time.Second)
	app.TPSCounter = 0
	app.FPSCounter = 0
}

func (app *LifeCycle) OnRender(elapsedTime time.Duration, mutex *sync.Mutex) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	now := time.Now()
	app.FPSCounter++
	if now.After(app.FPSPrinterTime) {
		fmt.Printf("%f FPS\n", app.FPSCounter/now.Sub(app.FPSStartTime).Seconds())
		app.FPSPrinterTime = now.Add(10 * time.Second)
	}
}

func (app *LifeCycle) OnTick(elapsedTime time.Duration, mutex *sync.Mutex) {
	now := time.Now()
	app.TPSCounter++
	if now.After(app.TPSPrinterTime) {
		fmt.Printf("%f TPS\n", app.TPSCounter/now.Sub(app.TPSStartTime).Seconds())
		app.TPSPrinterTime = now.Add(10 * time.Second)
	}
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
