package main

import (
	"fmt"
	"reflect"
	"syscall/js"
	// "net/http"
)

var (
	win           js.Value = js.Global()
	doc           js.Value = win.Get("document")
	body          js.Value = doc.Get("body")
	canvas        js.Value
	renderer      js.Func
	mousePosition Point
	laserPosition Point
)

/*
The js namespace is from https://golang.org/pkg/syscall/js
*/
func main() {
	runGameForever := make(chan bool) // https://stackoverflow.com/questions/47262088/golang-forever-channel
	fmt.Println("0", reflect.TypeOf(renderer))

	setupCanvas()

	renderer := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// canvasCtx := canvas.Call("getContext", "2d")
		// touchClickEventHandler := js.NewCallback(func(args []js.Value)) {
		// }
		// doc.Call("addEventListener", "click", touchClickEventHandler, false)
		// doc.Call("addEventListener", "touchstart", false)
		// http.Get("http://google.com")
		fmt.Println("Rendering!")
		fmt.Println("2", reflect.TypeOf(renderer))
		win.Call("requestAnimationFrame", renderer)
		return this
	})
	// postpones execution at the end; clean up memory
	defer renderer.Release()

	fmt.Println("1", reflect.TypeOf(renderer))

	// for the 60fps anims
	win.Call("requestAnimationFrame", renderer)

	// var mouseEventHandler js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 	// fmt.Println(this.Get("clientX").Float())
	// 	// fmt.Println(e.Get("clientX").Float())
	// 	fmt.Println("click tap")
	// 	// mousePosition.x = e.Get("clientX").Float()
	// 	// mousePosition.y = e.Get("clientY").Float()
	// 	return this
	// })
	// defer mouseEventHandler.Release()
	// win.Call("addEventListener", "click", mouseEventHandler, false)
	// win.Call("addEventListener", "touchstart", mouseEventHandler, false)

	<-runGameForever
}

// Helpers
type Point struct {
	x float64
	y float64
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
