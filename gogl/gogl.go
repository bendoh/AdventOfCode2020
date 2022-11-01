package gogl

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"image"
	"image/color"
	"log"
	"runtime"
)

const (
	vertexShaderSource = `
		#version 410
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 410
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(1, 1, 1, 1.0);
		}
	` + "\x00"
)

var Speed = 50

var width int32 = 640
var height int32 = 480

var window *glfw.Window

func initGlfw() {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	win, err := glfw.CreateWindow(int(width), int(height), "Advent of Code 2020 in GO", nil, nil)
	if err != nil {
		panic(err)
	}

	window = win
	window.MakeContextCurrent()
}

var program uint32
var pixels *image.RGBA
var texture, framebuffer uint32

func initOpenGL() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	pixels = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))

	gl.GenTextures(1, &texture)

	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

	gl.GenFramebuffers(1, &framebuffer)
	gl.BindFramebuffer(gl.FRAMEBUFFER, framebuffer)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, texture, 0)

	gl.BindFramebuffer(gl.READ_FRAMEBUFFER, framebuffer)
	gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, 0)
}

func InitWindow() *glfw.Window {
	initGlfw()
	initOpenGL()

	gl.UseProgram(program)

	return window
}

func Draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func DrawRect(c color.Color, x, y, w, h int) {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			pixels.Set(x+i, y+j, c)
		}
	}
}

func Poll() {
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels.Pix))

	gl.BlitFramebuffer(0, 0, width, height, 0, 0, width, height, gl.COLOR_BUFFER_BIT, gl.LINEAR)

	glfw.PollEvents()
	window.SwapBuffers()
}
