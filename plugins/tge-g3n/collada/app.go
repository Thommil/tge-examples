package main

import (
	bytes "bytes"
	fmt "fmt"
	time "time"

	tge "github.com/thommil/tge"

	camera "github.com/thommil/tge-g3n/camera"
	core "github.com/thommil/tge-g3n/core"
	gls "github.com/thommil/tge-g3n/gls"
	light "github.com/thommil/tge-g3n/light"
	collada "github.com/thommil/tge-g3n/loader/collada"
	math32 "github.com/thommil/tge-g3n/math32"
	renderer "github.com/thommil/tge-g3n/renderer"
)

type G3NApp struct {
	runtime  tge.Runtime
	gls      *gls.GLS
	scene    *core.Node
	camPersp *camera.Perspective
	renderer *renderer.Renderer
}

func (app *G3NApp) OnCreate(settings *tge.Settings) error {
	fmt.Println("OnCreate()")
	settings.Name = "G3NApp"
	settings.Fullscreen = true
	settings.EventMask = tge.AllEventsDisable
	return nil
}

func (app *G3NApp) OnStart(runtime tge.Runtime) error {
	fmt.Println("OnStart()")
	app.runtime = runtime

	runtime.Subscribe(tge.ResizeEvent{}.Channel(), app.OnResize)

	var err error

	// // Create OpenGL state
	app.gls, err = gls.New()
	if err != nil {
		return err
	}

	cc := math32.NewColor("black")
	app.gls.ClearColor(cc.R, cc.G, cc.B, 1)
	app.gls.Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)

	colladaStr, err := runtime.GetAsset("cyborg.dae")
	if err != nil {
		return err
	}

	reader := bytes.NewReader(colladaStr)

	dec, err := collada.DecodeReader(reader)
	if err != nil {
		fmt.Println(err)
		return err
	}

	app.scene = core.NewNode()
	decNode, err := dec.NewScene()
	if err != nil {
		fmt.Println(err)
		return err
	}
	app.renderer = renderer.NewRenderer(app.gls)
	app.renderer.SetScene(app.scene)
	err = app.renderer.AddDefaultShaders()
	if err != nil {
		return fmt.Errorf("Error from AddDefaulShaders:%v", err)
	}

	app.scene.Add(decNode)

	l1 := light.NewAmbient(&math32.Color{1, 1, 1}, 1.0)
	app.scene.Add(l1)

	// Add directional front  white light
	l3 := light.NewPoint(&math32.Color{1, 1, 1}, 1000.0)
	l3.SetPosition(10, 0, 10)
	app.scene.Add(l3)

	// Add directional front  white light
	l2 := light.NewPoint(&math32.Color{1, 1, 1}, 1000.0)
	l2.SetPosition(0, 20, 0)
	app.scene.Add(l2)

	app.camPersp = camera.NewPerspective(65, 1, 0.1, 100)
	app.camPersp.SetPosition(10, 10, 5)
	app.camPersp.LookAt(&math32.Vector3{0, 6, 0})

	return nil
}

func (app *G3NApp) OnResize(event tge.Event) bool {
	fmt.Printf("OnResize() : %v\n", event)
	app.camPersp.SetAspect(float32(event.(tge.ResizeEvent).Width) / float32(event.(tge.ResizeEvent).Height))
	app.gls.Viewport(0, 0, event.(tge.ResizeEvent).Width, event.(tge.ResizeEvent).Height)
	return false
}

func (app *G3NApp) OnResume() {
	fmt.Println("OnResume()")
}

func (app *G3NApp) OnRender(elapsedTime time.Duration, syncChan <-chan interface{}) {
	<-syncChan
	app.gls.Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
	app.renderer.Render(app.camPersp)
}

func (app *G3NApp) OnTick(elapsedTime time.Duration, syncChan chan<- interface{}) {
	syncChan <- true
}

func (app *G3NApp) OnPause() {
	fmt.Println("OnPause()")
}

func (app *G3NApp) OnStop() {
	fmt.Println("OnStop()")
}

func (app *G3NApp) OnDispose() {
	fmt.Println("OnDispose()")
}

func main() {
	tge.Run(&G3NApp{})
}
