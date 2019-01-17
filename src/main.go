package main

import (
	"math"
	"syscall/js" // https://golang.org/pkg/syscall/js
	// "net/http"
)

var (
	win              js.Value = js.Global()
	doc              js.Value = win.Get("document")
	body             js.Value = doc.Get("body")
	console          js.Value = win.Get("console")
	winSize          WinSize  = WinSize{w: 0, h: 0}
	canvas, laserCtx js.Value
	renderer         js.Func
	mouseXY, laserXY Point
)

func main() {
	// TODO explain this hack
	runGameForever := make(chan bool) // https://stackoverflow.com/questions/47262088/golang-forever-channel

	// http.Get("http://google.com")

	setupCanvas()
	setupGame()

	renderer = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		drawGame()
		win.Call("requestAnimationFrame", renderer)
		return nil
	})
	// postpones execution at the end; clean up memory
	defer renderer.Release()
	// for the 60fps anims
	win.Call("requestAnimationFrame", renderer)

	var mouseEventHandler js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		mouseXY.x = event.Get("clientX").Float()
		mouseXY.y = event.Get("clientY").Float()
		log("mouseEvent", "x", mouseXY.x, "y", mouseXY.y)
		log("isLaserCaught", isLaserCaught())
		return nil
	})
	defer mouseEventHandler.Release()
	win.Call("addEventListener", "click", mouseEventHandler, false)

	log("Game started")

	<-runGameForever
}

// Game helpers
func drawGame() {
	laserCtx.Call("clearRect", 0, 0, winSize.w, winSize.h)
	laserCtx.Call("beginPath")
	laserCtx.Call("arc", laserXY.x, laserXY.y, 20, 0, math.Pi*2)
	laserCtx.Set("fillStyle", "#0095DD")
	laserCtx.Call("fill")
	laserCtx.Call("closePath")

	laserXY.x += 0.02
	laserXY.y += -0.02
}

func setupGame() {
	laserCtx = canvas.Call("getContext", "2d")

	// center position
	laserXY.x = winSize.w / 2
	laserXY.y = winSize.h / 2
}

func isLaserCaught() bool {
	return laserCtx.Call("isPointInPath", mouseXY.x, mouseXY.y).Bool()
}

// Helpers
type Point struct {
	x, y float64
}

type WinSize struct {
	w, h float64
}

func setupCanvas() {
	winSize.h = win.Get("innerHeight").Float()
	winSize.w = win.Get("innerWidth").Float()

	canvas = doc.Call("createElement", "canvas")
	body.Call("appendChild", canvas)
	canvas.Set("height", winSize.h)
	canvas.Set("width", winSize.w)
}

func log(args ...interface{}) {
	console.Call("log", args...)
}
