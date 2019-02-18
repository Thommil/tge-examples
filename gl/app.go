package main

import (
	fmt "fmt"
	sync "sync"
	time "time"

	tge "github.com/thommil/tge"
	gl "github.com/thommil/tge-gl"
)

type GL struct {
	Runtime        tge.Runtime
	TPSStartTime   time.Time
	TPSPrinterTime time.Time
	FPSStartTime   time.Time
	FPSPrinterTime time.Time
	TPSCounter     float64
	FPSCounter     float64
	program        gl.Program
	vertexBuffer   gl.Buffer
	indexBuffer    gl.Buffer
}

func (app *GL) OnCreate(settings *tge.Settings) error {
	fmt.Println("OnCreate()")
	settings.Name = "GL"
	settings.Fullscreen = false
	settings.FPS = 100
	settings.TPS = 100

	return nil
}

func (app *GL) OnStart(runtime tge.Runtime) error {
	fmt.Println("OnStart()")
	app.Runtime = runtime

	runtime.Use(gl.GetPlugin())

	app.initProgram()
	app.initBuffers()

	return nil
}

func (app *GL) OnResize(width int, height int) {
	fmt.Printf("OnResize(%d, %d)\n", width, height)
	gl.Viewport(0, 0, width, height)
}

func (app *GL) OnResume() {
	fmt.Println("OnResume()")
	app.TPSStartTime = time.Now()
	app.FPSStartTime = time.Now()
	app.FPSPrinterTime = time.Now().Add(10 * time.Second)
	app.TPSPrinterTime = time.Now().Add(10 * time.Second)
	app.TPSCounter = 0
	app.FPSCounter = 0

	gl.ClearColor(0.15, 0.04, 0.15, 1)
}

func (app *GL) OnRender(elapsedTime time.Duration, mutex *sync.Mutex) {
	app.draw()
	now := time.Now()
	app.FPSCounter++
	if now.After(app.FPSPrinterTime) {
		fmt.Printf("%f FPS\n", app.FPSCounter/now.Sub(app.FPSStartTime).Seconds())
		app.FPSPrinterTime = now.Add(10 * time.Second)
	}
}

func (app *GL) OnTick(elapsedTime time.Duration, mutex *sync.Mutex) {
	now := time.Now()
	app.TPSCounter++
	if now.After(app.TPSPrinterTime) {
		fmt.Printf("%f TPS\n", app.TPSCounter/now.Sub(app.TPSStartTime).Seconds())
		app.TPSPrinterTime = now.Add(10 * time.Second)
	}
}

func (app *GL) initProgram() {
	fmt.Println("initProgram()")
	//// Shaders ////

	// Vertex shader source code
	vertCode := `
	attribute vec2 coordinates;
		
	void main(void) {
		gl_Position = vec4(coordinates, 0.0, 1.0);
	}`

	// Create a vertex shader object
	vertShader := gl.CreateShader(gl.VERTEX_SHADER)

	// Attach vertex shader source code
	gl.ShaderSource(vertShader, vertCode)

	// Compile the vertex shader
	gl.CompileShader(vertShader)

	//fragment shader source code
	fragCode := `
	void main(void) {
		gl_FragColor = vec4(1.0, 1.0, 1.0, 1.0);
	}`

	// Create fragment shader object
	fragShader := gl.CreateShader(gl.FRAGMENT_SHADER)

	// Attach fragment shader source code
	gl.ShaderSource(fragShader, fragCode)

	// Compile the fragmentt shader
	gl.CompileShader(fragShader)

	// Create a shader program object to store
	// the combined shader program
	app.program = gl.CreateProgram()

	// Attach a vertex shader
	gl.AttachShader(app.program, vertShader)

	// Attach a fragment shader
	gl.AttachShader(app.program, fragShader)

	// Link both the programs
	gl.LinkProgram(app.program)

	// Use the combined shader program object
	gl.UseProgram(app.program)

}

func (app *GL) initBuffers() {
	fmt.Println("initBuffers()")
	//// VERTEX BUFFER ////
	var vertices = []float32{
		// -0.5, 0.5,
		// -0.5, -0.5,
		// 0.5, -0.5,
		-1.0, -1.0,
		1.0, -1.0,
		0.0, 1.0,
	}

	// Create buffer
	app.vertexBuffer = gl.CreateBuffer()

	// Bind to buffer
	gl.BindBuffer(gl.ARRAY_BUFFER, app.vertexBuffer)

	// Pass data to buffer

	gl.BufferData(gl.ARRAY_BUFFER, gl.Float32ToBytes(vertices), gl.STATIC_DRAW)

	// Unbind buffer
	gl.BindBuffer(gl.ARRAY_BUFFER, gl.BufferNone)

	//// INDEX BUFFER ////
	var indices = []byte{
		2, 1, 0,
	}

	// Create buffer
	app.indexBuffer = gl.CreateBuffer()

	// Bind to buffer
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, app.indexBuffer)

	// Pass data to buffer
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, indices, gl.STATIC_DRAW)

	// Unbind buffer
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, gl.BufferNone)
}

func (app *GL) draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Bind vertex buffer object
	gl.BindBuffer(gl.ARRAY_BUFFER, app.vertexBuffer)

	// Bind index buffer object
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, app.indexBuffer)

	// Get the attribute location
	coord := gl.GetAttribLocation(app.program, "coordinates")

	// Enable the attribute
	gl.EnableVertexAttribArray(coord)

	// Point an attribute to the currently bound VBO
	gl.VertexAttribPointer(coord, 2, gl.FLOAT, false, 0, 0)

	//// Drawing the triangle ////
	// Draw the triangle
	gl.DrawElements(gl.TRIANGLES, 3, gl.UNSIGNED_BYTE, 0)
}

func (app *GL) OnPause() {
	fmt.Println("OnPause()")
}

func (app *GL) OnStop() {
	fmt.Println("OnStop()")
}

func (app *GL) OnDispose() error {
	fmt.Println("OnDispose()")
	return nil
}

func main() {
	tge.Run(&GL{})
}
