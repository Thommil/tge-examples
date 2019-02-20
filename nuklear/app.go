package main

import (
	fmt "fmt"
	sync "sync"
	time "time"

	tge "github.com/thommil/tge"
	gl "github.com/thommil/tge-gl"
	nk "github.com/thommil/tge-nuklear"
)

type NuklearApp struct {
	runtime tge.Runtime
	ctx     *nk.Context
}

func (app *NuklearApp) OnCreate(settings *tge.Settings) error {
	fmt.Println("OnCreate()")
	settings.Name = "NuklearApp"
	return nil
}

func (app *NuklearApp) OnStart(runtime tge.Runtime) error {
	fmt.Println("OnStart()")
	app.runtime = runtime
	runtime.Use(gl.GetPlugin())
	runtime.Use(nk.GetPlugin())

	app.ctx = nk.NkPlatformInit(nk.PlatformInstallCallbacks)
	atlas := nk.NewFontAtlas()
	nk.NkFontStashBegin(&atlas)
	sansFont := nk.NkFontAtlasAddDefault(atlas, 16, nil)
	nk.NkFontStashEnd()
	if sansFont != nil {
		nk.NkStyleSetFont(app.ctx, sansFont.Handle())
	}
	return nil
}

func (app *NuklearApp) OnResize(width int, height int) {
	fmt.Printf("OnResize(%d, %d)\n", width, height)
	gl.Viewport(0, 0, width, height)
}

func (app *NuklearApp) OnResume() {
	fmt.Println("OnResume()")
}

func (app *NuklearApp) OnRender(elapsedTime time.Duration, mutex *sync.Mutex) {
	nk.NkPlatformNewFrame()
	nk.NkClear(app.ctx)
	// Layout
	bounds := nk.NkRect(50, 50, 230, 250)
	update := nk.NkBegin(app.ctx, "Demo", bounds,
		nk.WindowBorder|nk.WindowMovable|nk.WindowScalable|nk.WindowMinimizable|nk.WindowTitle)

	if update > 0 {
		// nk.NkLayoutRowStatic(app.ctx, 30, 80, 1)
		// {
		// 	if nk.NkButtonLabel(app.ctx, "button") > 0 {
		// 		log.Println("[INFO] button pressed!")
		// 	}
		// }
		// nk.NkLayoutRowDynamic(app.ctx, 30, 2)
		// {
		// 	if nk.NkOptionLabel(app.ctx, "easy", flag(state.opt == Easy)) > 0 {
		// 		state.opt = Easy
		// 	}
		// 	if nk.NkOptionLabel(app.ctx, "hard", flag(state.opt == Hard)) > 0 {
		// 		state.opt = Hard
		// 	}
		// }
		// nk.NkLayoutRowDynamic(app.ctx, 25, 1)
		// {
		// 	nk.NkPropertyInt(app.ctx, "Compression:", 0, &state.prop, 100, 10, 1)
		// }
		// nk.NkLayoutRowDynamic(app.ctx, 20, 1)
		// {
		// 	nk.NkLabel(app.ctx, "background:", nk.TextLeft)
		// }
		// nk.NkLayoutRowDynamic(app.ctx, 25, 1)
		// {
		// 	size := nk.NkVec2(nk.NkWidgetWidth(app.ctx), 400)
		// 	if nk.NkComboBeginColor(app.ctx, state.bgColor, size) > 0 {
		// 		nk.NkLayoutRowDynamic(app.ctx, 120, 1)
		// 		state.bgColor = nk.NkColorPicker(app.ctx, state.bgColor, nk.ColorFormatRGBA)
		// 		nk.NkLayoutRowDynamic(app.ctx, 25, 1)
		// 		r, g, b, a := state.bgColor.RGBAi()
		// 		r = nk.NkPropertyi(app.ctx, "#R:", 0, r, 255, 1, 1)
		// 		g = nk.NkPropertyi(app.ctx, "#G:", 0, g, 255, 1, 1)
		// 		b = nk.NkPropertyi(app.ctx, "#B:", 0, b, 255, 1, 1)
		// 		a = nk.NkPropertyi(app.ctx, "#A:", 0, a, 255, 1, 1)
		// 		state.bgColor.SetRGBAi(r, g, b, a)
		// 		nk.NkComboEnd(app.ctx)
		// 	}
		// }
	}
	nk.NkEnd(app.ctx)

	// // Render
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.ClearColor(0, 0, 0, 1)
	// nk.NkPlatformRender(nk.AntiAliasingOn, maxVertexBuffer, maxElementBuffer)
}

func (app *NuklearApp) OnTick(elapsedTime time.Duration, mutex *sync.Mutex) {

}

func (app *NuklearApp) OnMouseEvent(event tge.MouseEvent) {
	// NOP
}

func (app *NuklearApp) OnScrollEvent(event tge.ScrollEvent) {
	// NOP
}

func (app *NuklearApp) OnKeyEvent(event tge.KeyEvent) {
	// NOP
}

func (app *NuklearApp) OnPause() {
	fmt.Println("OnPause()")
}

func (app *NuklearApp) OnStop() {
	fmt.Println("OnStop()")
}

func (app *NuklearApp) OnDispose() error {
	fmt.Println("OnDispose()")
	return nil
}

func main() {
	tge.Run(&NuklearApp{})
}
