package main

import (
	"math"
	"syscall/js" // https://golang.org/pkg/syscall/js
)

var (
	window                       js.Value = js.Global()
	document                     js.Value = window.Get("document")
	body                         js.Value = document.Get("body")
	windowSize                   WindowSize
	canvas, laserCtx             js.Value
	mousePosition, laserPosition Point
	renderer                     js.Func
	dx, dy                       float64  = 2.5, -2.5
	ballRadius                   float64  = 35
	beep                         js.Value = document.Call("getElementById", "beep")
)

func main() {
	runGameForever := make(chan bool) // explain TODO https://stackoverflow.com/questions/47262088/golang-forever-channel

	setup()

	renderer = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		updateGame()
		window.Call("requestAnimationFrame", renderer)
		return nil
	})
	defer renderer.Release()                       // postpones execution at the end; clean up memory
	window.Call("requestAnimationFrame", renderer) // for the 60fps anims

	var mouseEventHandler js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		updateGame()
		updatePlayer(args[0])
		return nil
	})
	defer mouseEventHandler.Release()
	window.Call("addEventListener", "click", mouseEventHandler, false)

	<-runGameForever
}

func updateGame() {
	// wall detection
	if laserPosition.x+dx > windowSize.w-ballRadius || laserPosition.x+dx < ballRadius {
		dx = -dx
	}

	if laserPosition.y+dy > windowSize.h-ballRadius || laserPosition.y+dy < ballRadius {
		dy = -dy
	}

	laserPosition.x += dx
	laserPosition.y += dy

	laserCtx.Call("clearRect", 0, 0, windowSize.w, windowSize.h)
	laserCtx.Call("beginPath")
	laserCtx.Call("arc", laserPosition.x, laserPosition.y, ballRadius, 0, math.Pi*2)
	laserCtx.Call("fill")
	laserCtx.Call("closePath")
}

func updatePlayer(event js.Value) {
	mousePosition.x = event.Get("clientX").Float()
	mousePosition.y = event.Get("clientY").Float()
	log("mouseEvent", "x", mousePosition.x, "y", mousePosition.y)

	if isLaserCaught() {
		playSound() // figure out the delay
		blinkLamp()
	}
}

// Helpers
func setup() {
	windowSize.h = window.Get("innerHeight").Float()
	windowSize.w = window.Get("innerWidth").Float()

	canvas = document.Call("createElement", "canvas")
	body.Call("appendChild", canvas)
	canvas.Set("height", windowSize.h)
	canvas.Set("width", windowSize.w)

	laserCtx = canvas.Call("getContext", "2d")
	laserCtx.Set("fillStyle", "red")
	laserPosition.x = windowSize.w / 2
	laserPosition.y = windowSize.h / 2
}

func isLaserCaught() bool {
	return laserCtx.Call("isPointInPath", mousePosition.x, mousePosition.y).Bool() // TODO fix hitmark
}

func playSound() {
	beep.Call("play") // there is a delay when running chrome on android
	window.Get("navigator").Call("vibrate", 300)
}

func blinkLamp() {
	// http.Get("http://192.168.31.123:8080/blink/twice") // excluded since no tree shaking, made it 7MB
}

type Point struct {
	x, y float64
}

type WindowSize struct {
	w, h float64
}

func log(args ...interface{}) {
	window.Get("console").Call("log", args...)
}
