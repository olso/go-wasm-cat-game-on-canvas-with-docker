package main

import (
	"runtime"
	"syscall/js"
	// "net/http"
)

type RedDotPosition struct {
	X     float64
	Y     float64
	id    int
	alive bool
}

var (
	win        js.Value = js.Global()
	doc        js.Value = win.Get("document")
	body       js.Value = doc.Get("body")
	canvas     js.Value
	renderGame js.Func
)

/*
The js namespace is from https://golang.org/pkg/syscall/js
*/
func game() {
	setupCanvas()

	renderGame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// canvasCtx := canvas.Call("getContext", "2d")
		// touchClickEventHandler := js.NewCallback(func(args []js.Value)) {
		// }
		// doc.Call("addEventListener", "click", touchClickEventHandler, false)
		// doc.Call("addEventListener", "touchstart", false)
		// http.Get("http://google.com")
		// fmt.Println("Rendering!")
		win.Call("requestAnimationFrame", renderGame)
		return nil
	})

	// postpones execution at the end; clean up memory
	defer renderGame.Release()

	// prefer browser repaint in next event loop, think 60fps anims
	win.Call("requestAnimationFrame", renderGame)
}

func main() {
	go game()        // Run game in goroutine, run forever
	runtime.Goexit() // https://golang.org/pkg/runtime/#Goexit
}

func setupCanvas() {
	canvas = doc.Call("createElement", "canvas")
	body.Call("appendChild", canvas)
	winHeight, winWidth := getWinSize()
	canvas.Set("height", winHeight)
	canvas.Set("width", winWidth)
}

func getWinSize() (float64, float64) {
	h := win.Get("innerHeight").Float()
	w := win.Get("innerWidth").Float()

	return h, w
}
