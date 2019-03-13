package main

import (
	"fmt"
	"time"

	tge "github.com/thommil/tge"
	audio "github.com/thommil/tge-audio"
)

type AudioApp struct {
	runtime         tge.Runtime
	audioInit       bool
	sampleData      []byte
	sampleBuffer    audio.Buffer
	stereoPanNode   audio.StereoPannerNode
	gainNode        audio.GainNode
	destinationNode audio.DestinationNode
	width           int32
	height          int32
}

func (app *AudioApp) OnCreate(settings *tge.Settings) error {
	fmt.Println("OnCreate()")
	settings.Fullscreen = true
	return nil
}

func (app *AudioApp) OnStart(runtime tge.Runtime) error {
	fmt.Println("OnStart()")
	app.runtime = runtime

	runtime.Subscribe(tge.MouseEvent{}.Channel(), app.OnMouseEvent)
	runtime.Subscribe(tge.ResizeEvent{}.Channel(), app.OnResize)
	return nil
}

// InitAudio is start on mouse click only due to browsers restrictions
func (app *AudioApp) InitAudio() error {
	var err error

	//Buffer (should be done in a loading screen off course)
	fmt.Println("Loading bass.wav in buffer")
	if app.sampleBuffer, err = audio.CreateBuffer("bass.wav"); err != nil {
		return err
	}

	// Audio graph
	fmt.Println("Creating audio graph")
	if app.destinationNode, err = audio.CreateDestinationNode(); err != nil {
		return err
	}
	if app.gainNode, err = audio.CreateGainNode(); err != nil {
		return err
	}
	if app.stereoPanNode, err = audio.CreateStereoPannerNode(); err != nil {
		return err
	}
	// destination is connected at resume state only
	app.stereoPanNode.Connect(app.gainNode)
	app.audioInit = true

	app.gainNode.Connect(app.destinationNode)

	return nil
}

func (app *AudioApp) OnResume() {
	fmt.Println("OnResume()")
	// Open sound
	if app.audioInit {
		app.gainNode.Connect(app.destinationNode)
	}
}

func (app *AudioApp) OnResize(event tge.Event) bool {
	app.width = (event.(tge.ResizeEvent)).Width
	app.height = (event.(tge.ResizeEvent)).Height
	return false
}

func (app *AudioApp) OnMouseEvent(event tge.Event) bool {
	mouseEvent := event.(tge.MouseEvent)
	if mouseEvent.Type == tge.TypeDown && (mouseEvent.Button&tge.ButtonLeftOrTouchFirst != 0) {
		if !app.audioInit {
			if err := app.InitAudio(); err != nil {
				fmt.Printf("ERROR: %s\n", err)
			}
		}

		app.stereoPanNode.Pan(float32(mouseEvent.X)/float32(app.width)*2 - 1.0)
		app.gainNode.Gain(float32(app.height-mouseEvent.Y) / float32(app.height))

		fmt.Printf("Set pan to %v and gain to %v%%\n", (float32(mouseEvent.X)/float32(app.width)*2 - 1.0), float32(app.height-mouseEvent.Y)/float32(app.height)*100)

		if sourceNode, err := audio.CreateBufferSourceNode(app.sampleBuffer); err != nil {
			fmt.Printf("ERROR: %s\n", err)
		} else {
			sourceNode.Connect(app.stereoPanNode)
			sourceNode.Start(false, 0, 0)
		}
	}
	return false
}

func (app *AudioApp) OnRender(elaspedTime time.Duration, syncChan <-chan interface{}) {
	<-syncChan
}

func (app *AudioApp) OnTick(elaspedTime time.Duration, syncChan chan<- interface{}) {
	syncChan <- true
}

func (app *AudioApp) OnPause() {
	fmt.Println("OnPause()")
	// Close sound
	app.gainNode.Disconnect(app.destinationNode)
}

func (app *AudioApp) OnStop() {
	fmt.Println("OnStop()")
	app.runtime.Unsubscribe(tge.MouseEvent{}.Channel(), app.OnMouseEvent)
	app.runtime.Unsubscribe(tge.ResizeEvent{}.Channel(), app.OnResize)
}

func (app *AudioApp) OnDispose() {
	fmt.Println("OnDispose()")
}

func main() {
	tge.Run(&AudioApp{})
}
