package main

import (
	"bytes"
	fmt "fmt"
	"image/jpeg"
	sync "sync"
	time "time"

	tge "github.com/thommil/tge"
)

type AssetsApp struct {
	Runtime tge.Runtime
}

func (app *AssetsApp) OnCreate(settings *tge.Settings) error {
	fmt.Println("OnCreate()")
	settings.Name = "AssetsAppApp"
	settings.Fullscreen = false
	settings.FPS = 10
	settings.TPS = 10
	settings.EventMask = tge.AllEventsDisable
	return nil
}

func (app *AssetsApp) OnStart(runtime tge.Runtime) error {
	fmt.Println("OnStart()")
	app.Runtime = runtime
	return nil
}

func (app *AssetsApp) OnResize(width int, height int) {
	fmt.Printf("OnResize(%d, %d)\n", width, height)
}

func (app *AssetsApp) OnResume() {
	fmt.Println("OnResume()")
	if txtContent, err := app.Runtime.LoadAsset("test.txt"); err != nil {
		fmt.Printf("Error loading TXT file : %s\n", err)
	} else {
		fmt.Printf("test.txt : %s\n", string(txtContent))
	}

	if jpgContent, err := app.Runtime.LoadAsset("test.jpg"); err != nil {
		fmt.Printf("Error loading JPG file : %s\n", err)
	} else {
		if img, err := jpeg.Decode(bytes.NewBuffer(jpgContent)); err != nil {
			fmt.Printf("Error dcoding JPG file : %s\n", err)
		} else {
			bounds := img.Bounds()
			var histogram [16][4]int
			for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					r, g, b, a := img.At(x, y).RGBA()
					// A color's RGBA method returns values in the range [0, 65535].
					// Shifting by 12 reduces this to the range [0, 15].
					histogram[r>>12][0]++
					histogram[g>>12][1]++
					histogram[b>>12][2]++
					histogram[a>>12][3]++
				}
			}

			// Print the results.
			fmt.Printf("%-14s %6s %6s %6s %6s\n", "bin", "red", "green", "blue", "alpha")
			for i, x := range histogram {
				fmt.Printf("0x%04x-0x%04x: %6d %6d %6d %6d\n", i<<12, (i+1)<<12-1, x[0], x[1], x[2], x[3])
			}
		}
	}

}

func (app *AssetsApp) OnRender(elapsedTime time.Duration, mutex *sync.Mutex) {
	fmt.Println("OnRender()")
}

func (app *AssetsApp) OnTick(elapsedTime time.Duration, mutex *sync.Mutex) {
	fmt.Println("OnTick()")

}

func (app *AssetsApp) OnMouseEvent(event tge.MouseEvent) {
	// NOP
}

func (app *AssetsApp) OnScrollEvent(event tge.ScrollEvent) {
	// NOP
}

func (app *AssetsApp) OnKeyEvent(event tge.KeyEvent) {
	// NOP
}

func (app *AssetsApp) OnPause() {
	fmt.Println("OnPause()")
}

func (app *AssetsApp) OnStop() {
	fmt.Println("OnStop()")
}

func (app *AssetsApp) OnDispose() error {
	fmt.Println("OnDispose()")
	return nil
}

func main() {
	tge.Run(&AssetsApp{})
}
