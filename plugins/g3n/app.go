package main

import (
	"bytes"
	fmt "fmt"
	sync "sync"
	time "time"

	"github.com/thommil/tge-g3n/light"
	"github.com/thommil/tge-g3n/loader/collada"

	control "github.com/thommil/tge-g3n/camera/control"

	tge "github.com/thommil/tge"
	g3n "github.com/thommil/tge-g3n"
	camera "github.com/thommil/tge-g3n/camera"
	core "github.com/thommil/tge-g3n/core"
	gls "github.com/thommil/tge-g3n/gls"
	math32 "github.com/thommil/tge-g3n/math32"
	renderer "github.com/thommil/tge-g3n/renderer"
)

type G3NApp struct {
	runtime     tge.Runtime
	gls         *gls.GLS
	scene       *core.Node
	camPersp    *camera.Perspective
	renderer    *renderer.Renderer
	orbCtrl     *control.OrbitControl
	animTargets map[string]*collada.AnimationTarget

	TPSStartTime   time.Time
	TPSPrinterTime time.Time
	FPSStartTime   time.Time
	FPSPrinterTime time.Time
	TPSCounter     float64
	FPSCounter     float64
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
	runtime.Use(g3n.GetPlugin())

	app.TPSStartTime = time.Now()
	app.FPSStartTime = time.Now()
	app.FPSPrinterTime = time.Now().Add(10 * time.Second)
	app.TPSPrinterTime = time.Now().Add(10 * time.Second)
	app.TPSCounter = 0
	app.FPSCounter = 0

	var err error

	// // Create OpenGL state
	app.gls, err = gls.New()
	if err != nil {
		return err
	}

	cc := math32.NewColor("black")
	app.gls.ClearColor(cc.R, cc.G, cc.B, 1)
	app.gls.Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)

	colladaStr, err := runtime.LoadAsset("cyborg.dae")
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

	// Checks for animations
	app.animTargets, err = dec.NewAnimationTargets(decNode)
	if err == nil {
		for _, at := range app.animTargets {
			at.SetStart(-1.0)
			at.Reset()
			at.SetLoop(true)
		}
	}

	// app.scene = core.NewNode()
	// app.renderer = renderer.NewRenderer(app.gls)
	// err = app.renderer.AddDefaultShaders()
	// if err != nil {
	// 	return fmt.Errorf("Error from AddDefaulShaders:%v", err)
	// }
	// app.renderer.SetScene(app.scene)

	// TORUS
	// geom := geometry.NewTorus(1, .4, 480, 1280, math32.Pi*2)
	// mat := material.NewPhong(&math32.Color{0.15, 0.04, 0.15})
	// torusMesh := graphic.NewMesh(geom, mat)
	// app.scene.Add(torusMesh)

	// PERFS
	// torusGeometry := geometry.NewTorus(0.5, 0.2, 16, 16, 2*math32.Pi)
	// halfSize := 5
	// step := 2
	// count := 0
	// for i := -halfSize; i < (halfSize + 1); i += step {
	// 	for j := -halfSize; j < (halfSize + 1); j += step {
	// 		for k := -halfSize; k < (halfSize + 1); k += step {
	// 			count += 1
	// 			mat := material.NewStandard(&math32.Color{rand.Float32(), rand.Float32(), rand.Float32()})
	// 			//mat.SetSpecularColor(math32.NewColor("white"))
	// 			//mat.SetShininess(100)
	// 			torus := graphic.NewMesh(torusGeometry, mat)
	// 			torus.SetPosition(float32(i), float32(j), float32(k))
	// 			torus.SetRotation(rand.Float32()*2*math32.Pi, rand.Float32()*2*math32.Pi, rand.Float32()*2*math32.Pi)
	// 			app.scene.Add(torus)
	// 		}
	// 	}
	// }

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

/*
func (t *LoaderCollada) load(a *app.App, path string) error {

	// Remove previous model from the scene
	if t.prevLoaded != nil {
		a.Scene().Remove(t.prevLoaded)
		t.prevLoaded.Dispose()
		t.prevLoaded = nil
	}

	// Decodes collada file
	dec, err := collada.Decode(path)
	if err != nil && err != io.EOF {
		t.selFile.SetError(err.Error())
		return err
	}
	dec.SetDirImages(a.DirData() + "/images")

	// Loads collada scene
	s, err := dec.NewScene()
	if err != nil {
		t.selFile.SetError(err.Error())
		return err
	}
	a.Scene().Add(s)
	t.prevLoaded = s

	// Checks for animations
	ats, err := dec.NewAnimationTargets(s)
	if err == nil {
		t.animTargets = ats
		for _, at := range ats {
			at.SetStart(-1.0)
			at.Reset()
			at.SetLoop(true)
		}
	}
	return nil
}

func (t *LoaderCollada) Render(a *app.App) {

	if t.animTargets != nil {
		dt := a.FrameDeltaSeconds()
		for _, at := range t.animTargets {
			at.Update(dt)
		}
	}
}


*/

func (app *G3NApp) OnResize(width int, height int) {
	fmt.Printf("OnResize(%d, %d)\n", width, height)
	app.camPersp.SetAspect(float32(width) / float32(height))
}

func (app *G3NApp) OnResume() {
	fmt.Println("OnResume()")
}

func (app *G3NApp) OnRender(elapsedTime time.Duration, mutex *sync.Mutex) {
	mutex.Lock()
	_, err := app.renderer.Render(app.camPersp)
	mutex.Unlock()

	if err != nil {
		fmt.Printf("%v\n", err)
		app.runtime.Stop()
	}

	now := time.Now()
	app.FPSCounter++
	if now.After(app.FPSPrinterTime) {
		fmt.Printf("%f FPS\n", app.FPSCounter/now.Sub(app.FPSStartTime).Seconds())
		app.FPSPrinterTime = now.Add(10 * time.Second)
	}

}

func (app *G3NApp) OnTick(elapsedTime time.Duration, mutex *sync.Mutex) {
	now := time.Now()
	app.TPSCounter++
	if now.After(app.TPSPrinterTime) {
		fmt.Printf("%f TPS\n", app.TPSCounter/now.Sub(app.TPSStartTime).Seconds())
		app.TPSPrinterTime = now.Add(10 * time.Second)
	}
	// mutex.Lock()
	// for _, node := range app.scene.Children() {
	// 	switch node.(type) {
	// 	case *graphic.Mesh:
	// 		node.(*graphic.Mesh).RotateY(0.05)
	// 	}
	// }

	// //app.scene.ChildAt(0).(*graphic.Mesh).RotateX(0.01)
	// mutex.Unlock()

	mutex.Lock()
	app.scene.RotateY(0.01)
	mutex.Unlock()
}

var mouseDown bool
var lastMoveEvent tge.MouseEvent

func (app *G3NApp) OnMouseEvent(event tge.MouseEvent) {
	switch event.Type {
	case tge.TypeDown:
		mouseDown = true
	case tge.TypeUp:
		mouseDown = false
		lastMoveEvent.X = 0
		lastMoveEvent.Y = 0
	case tge.TypeMove:
		if mouseDown {
			if lastMoveEvent.X != 0 {
				app.orbCtrl.RotateUp(float32(event.Y-lastMoveEvent.Y) * 0.01)
				app.orbCtrl.RotateLeft(float32(event.X-lastMoveEvent.X) * 0.01)
			}
			lastMoveEvent = event
		}
	}
}

func (app *G3NApp) OnScrollEvent(event tge.ScrollEvent) {
	app.orbCtrl.Zoom(float32(-event.Y) * 0.5)
}

func (app *G3NApp) OnKeyEvent(event tge.KeyEvent) {
	// NOP
}

func (app *G3NApp) OnPause() {
	fmt.Println("OnPause()")
}

func (app *G3NApp) OnStop() {
	fmt.Println("OnStop()")
}

func (app *G3NApp) OnDispose() error {
	fmt.Println("OnDispose()")
	return nil
}

func main() {
	tge.Run(&G3NApp{})
}
