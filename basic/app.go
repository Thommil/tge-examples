package main

import (
	log "log"

	tge "github.com/thommil/tge"
	physics "github.com/thommil/tge/physics"
	player "github.com/thommil/tge/player"
	renderer "github.com/thommil/tge/renderer"
	ui "github.com/thommil/tge/ui"
)

type BasicApp struct {
}

func (app BasicApp) OnCreate(settings *tge.Settings) error {
	log.Println("Create()")
	settings.Name = "BasicApp"
	settings.Fullscreen = true
	return nil
}

func (app BasicApp) OnStart(runtime tge.Runtime) error {
	log.Println("Start()")
	return nil
}

func (app BasicApp) OnResize(width int, height int) {
	log.Printf("Resize(%d, %d)\n", width, height)
}

func (app BasicApp) OnResume() {
	log.Println("Resume()")
}

func (app BasicApp) OnRender(renderer renderer.Renderer, ui ui.UI, player player.Player) {
	log.Println("Render()")
}

func (app BasicApp) OnTick(physics physics.Physics) {
	log.Println("Tick()")
}

func (app BasicApp) OnPause() {
	log.Println("Pause()")
}

func (app BasicApp) OnStop() {
	log.Println("Stop()")
}

func (app BasicApp) OnDispose() error {
	log.Println("Dispose()")
	return nil
}

func main() {
	app := BasicApp{}
	tge.Run(app)
}
