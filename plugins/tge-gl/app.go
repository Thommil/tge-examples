package main

import (
	fmt "fmt"
	time "time"

	tge "github.com/thommil/tge"
	gl "github.com/thommil/tge-gl"
)

type GLApp struct {
	runtime      tge.Runtime
	program      gl.Program
	vertexBuffer gl.Buffer
	vao          gl.VertexArray
	indexBuffer  gl.Buffer
}

func (app *GLApp) OnCreate(settings *tge.Settings) error {
	fmt.Println("OnCreate()")
	settings.Name = "GLApp"
	settings.Fullscreen = false
	settings.EventMask = tge.AllEventsDisable
	return nil
}

func (app *GLApp) OnStart(runtime tge.Runtime) error {
	fmt.Println("OnStart()")
	app.runtime = runtime

	runtime.Subscribe(tge.ResizeEvent{}.Channel(), app.OnResize)

	gl.ClearColor(0.15, 0.04, 0.15, 1)
	app.initProgram()
	app.initBuffers()

	return nil
}

func (app *GLApp) OnResize(event tge.Event) bool {
	gl.Viewport(0, 0, int(event.(tge.ResizeEvent).Width), int(event.(tge.ResizeEvent).Height))
	return false
}

func (app *GLApp) OnResume() {
	fmt.Println("OnResume()")
}

func (app *GLApp) OnRender(elapsedTime time.Duration, syncChan <-chan interface{}) {
	<-syncChan
	app.draw()
}

func (app *GLApp) OnTick(elapsedTime time.Duration, syncChan chan<- interface{}) {
	syncChan <- true
}

func (app *GLApp) initProgram() {
	fmt.Println("initProgram()")
	//// Shaders ////

	// Vertex shader source code
	vertCode := fmt.Sprintf("#version %s\n", gl.GetGLSLVersion())
	vertCode += `layout(location = 0) in vec2 coordinates;

	void main() {
		gl_Position = vec4(coordinates, 0.0, 1.0);
	}`

	// Create a vertex shader object
	vertShader := gl.CreateShader(gl.VERTEX_SHADER)

	// Attach vertex shader source code
	gl.ShaderSource(vertShader, vertCode)

	// Compile the vertex shader
	gl.CompileShader(vertShader)

	//fragment shader source code
	fragCode := fmt.Sprintf("#version %s\n", gl.GetGLSLVersion())
	fragCode += `precision mediump float;

	out vec4 FragColor;

	void main() {
		FragColor = vec4(0.85, 0.8, 0.8, 1.0);
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

func (app *GLApp) initBuffers() {
	fmt.Println("initBuffers()")
	//// VERTEX BUFFER ////
	var vertices = []float32{
		-1.0, -1.0,
		1.0, -1.0,
		0.0, 1.0,
	}

	// Create VAO
	app.vao = gl.CreateVertexArray()

	// Bin VAO
	gl.BindVertexArray(app.vao)

	// Create buffer
	app.vertexBuffer = gl.CreateBuffer()

	// Bind to buffer
	gl.BindBuffer(gl.ARRAY_BUFFER, app.vertexBuffer)

	// Pass data to buffer
	gl.BufferData(gl.ARRAY_BUFFER, gl.Float32ToBytes(vertices), gl.STATIC_DRAW)

	// Unbind buffer
	gl.BindBuffer(gl.ARRAY_BUFFER, gl.NONE)

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
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, gl.NONE)

}

func (app *GLApp) draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Bin VAO
	gl.BindVertexArray(app.vao)

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

	// Draw the triangle
	gl.DrawElements(gl.TRIANGLES, 3, gl.UNSIGNED_BYTE, 0)
}

func (app *GLApp) OnPause() {
	fmt.Println("OnPause()")
}

func (app *GLApp) OnStop() {
	fmt.Println("OnStop()")
	app.runtime.Unsubscribe(tge.ResizeEvent{}.Channel(), app.OnResize)
}

func (app *GLApp) OnDispose() {
	fmt.Println("OnDispose()")
}

func main() {
	tge.Run(&GLApp{})
}
