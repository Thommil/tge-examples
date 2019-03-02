package main

import (
	bytes "bytes"
	fmt "fmt"
	sync "sync"
	time "time"

	tge "github.com/thommil/tge"

	g3n "github.com/thommil/tge-g3n"

	camera "github.com/thommil/tge-g3n/camera"
	control "github.com/thommil/tge-g3n/camera/control"
	core "github.com/thommil/tge-g3n/core"
	gls "github.com/thommil/tge-g3n/gls"
	light "github.com/thommil/tge-g3n/light"
	collada "github.com/thommil/tge-g3n/loader/collada"
	math32 "github.com/thommil/tge-g3n/math32"
	renderer "github.com/thommil/tge-g3n/renderer"
)

type G3NApp struct {
	runtime    tge.Runtime
	gls        *gls.GLS
	scene      *core.Node
	camPersp   *camera.Perspective
	renderer   *renderer.Renderer
	orbCtrl    *control.OrbitControl
	orbCtrlMvt [3]int32
}

func (app *G3NApp) OnCreate(settings *tge.Settings) error {
	fmt.Println("OnCreate()")
	settings.Name = "G3NApp"
	settings.Fullscreen = true
	settings.TPS = 100
	settings.EventMask = tge.MouseMotionEventEnabled | tge.ScrollEventEnabled | tge.MouseButtonEventEnabled
	return nil
}

func (app *G3NApp) OnStart(runtime tge.Runtime) error {
	fmt.Println("OnStart()")
	app.runtime = runtime
	runtime.Use(g3n.GetInstance())

	runtime.Subscribe(tge.ResizeEvent{}.Channel(), app.OnResize)
	runtime.Subscribe(tge.MouseEvent{}.Channel(), app.OnMouseEvent)
	runtime.Subscribe(tge.ScrollEvent{}.Channel(), app.OnScrollEvent)

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

	app.camPersp = camera.NewPerspective(65, 1, 0.01, 1000)
	app.camPersp.SetPosition(10, 10, 5)
	app.camPersp.LookAt(&math32.Vector3{0, 6, 0})

	app.orbCtrl = control.NewOrbitControl(app.camPersp)

	return nil
}

func (app *G3NApp) OnResize(event tge.Event) bool {
	app.camPersp.SetAspect(float32(event.(tge.ResizeEvent).Width) / float32(event.(tge.ResizeEvent).Height))
	return false
}

func (app *G3NApp) OnResume() {
	fmt.Println("OnResume()")
}

func (app *G3NApp) OnRender(elapsedTime time.Duration, mutex *sync.Mutex) {
	mutex.Lock()
	app.gls.Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
	_, err := app.renderer.Render(app.camPersp)
	mutex.Unlock()

	if err != nil {
		fmt.Printf("%v\n", err)
		app.runtime.Stop()
	}
}

func (app *G3NApp) OnTick(elapsedTime time.Duration, mutex *sync.Mutex) {
	mutex.Lock()
	//app.scene.RotateY(0.01)
	app.orbCtrl.RotateUp(float32(app.orbCtrlMvt[0]) * 0.01)
	app.orbCtrl.RotateLeft(float32(app.orbCtrlMvt[1]) * 0.01)
	app.orbCtrl.Zoom(float32(app.orbCtrlMvt[2]) * 0.5)
	app.orbCtrlMvt[0] = 0
	app.orbCtrlMvt[1] = 0
	app.orbCtrlMvt[2] = 0
	mutex.Unlock()
}

var mouseDown bool
var lastMoveEvent tge.MouseEvent

func (app *G3NApp) OnMouseEvent(event tge.Event) bool {
	e := event.(tge.MouseEvent)
	switch e.Type {
	case tge.TypeDown:
		mouseDown = true
	case tge.TypeUp:
		mouseDown = false
		lastMoveEvent.X = 0
		lastMoveEvent.Y = 0
	case tge.TypeMove:
		if mouseDown {
			if lastMoveEvent.X != 0 || lastMoveEvent.Y != 0 {
				app.orbCtrlMvt[0] += e.Y - lastMoveEvent.Y
				app.orbCtrlMvt[1] += e.X - lastMoveEvent.X
			}
			lastMoveEvent = e
		}
	}
	return false
}

func (app *G3NApp) OnScrollEvent(event tge.Event) bool {
	e := event.(tge.ScrollEvent)
	app.orbCtrlMvt[2] += -e.Y
	return false
}

func (app *G3NApp) OnPause() {
	fmt.Println("OnPause()")
}

func (app *G3NApp) OnStop() {
	fmt.Println("OnStop()")
	app.runtime.Unsubscribe(tge.ResizeEvent{}.Channel(), app.OnResize)
	app.runtime.Unsubscribe(tge.MouseEvent{}.Channel(), app.OnMouseEvent)
	app.runtime.Unsubscribe(tge.ScrollEvent{}.Channel(), app.OnScrollEvent)
}

func (app *G3NApp) OnDispose() error {
	fmt.Println("OnDispose()")
	return nil
}

func main() {
	tge.Run(&G3NApp{})
}
