package main

import (
	"syscall/js"
	// "net/http"
)

// type Point struct {
// 	X   float64
// 	Y   float64
// 	Num int
// }

func main() {
	window := js.Global()
	document := window.Get("document")
	body := document.Get("body")
	canvas := document.Call("createElement", "canvas")
	body.Call("appendChild", canvas)

	winHeight := window.Get("innerHeight").Float()
	winWidth := window.Get("innerWidth").Float()
	canvas.Set("height", winHeight)
	canvas.Set("width", winWidth)

	// canvasCtx := canvas.Call("getContext", "2d")

	// touchClickEventHandler := js.NewCallback(func(args []js.Value)) {
	// }
	// doc.Call("addEventListener", "click", touchClickEventHandler, false)
	// doc.Call("addEventListener", "touchstart", false)

	// alert := js.Global().Get("alert")
	// alert.Invoke("Hello wasm!")
	// http.Get("http://google.com")
}
