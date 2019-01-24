package main

import (
	"math"
	"syscall/js" // https://golang.org/pkg/syscall/js
)

var (
	window, doc, body, canvas, laserCtx, beep js.Value
	windowSize                                struct{ w, h float64 }
	gs                                        = gameState{laserSize: 35, directionX: 3.7, directionY: -3.7, laserX: 40, laserY: 40}
)

func main() {
	// https://stackoverflow.com/a/47262117
	// creates empty channel
	runGameForever := make(chan bool)
	setup()

	// declare renderer at compile time
	var renderer js.Func
	// looks like JS callback, right 😌
	renderer = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		updateGame()
		// for the 60fps anims
		window.Call("requestAnimationFrame", renderer)
		return nil
	})
	window.Call("requestAnimationFrame", renderer)

	var mouseEventHandler js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		updatePlayer(args[0])
		return nil
	})
	window.Call("addEventListener", "pointerdown", mouseEventHandler)

	// attempt to receive from empty channel
	// since noone ever sends anything on it, it's essentially a blocking forever operation
	// we basically have a daeomon/service/background program
	// in WASM world, our game will keep running 😉
	<-runGameForever
}

func updateGame() {
	if gs.laserX+gs.directionX > windowSize.w-gs.laserSize || gs.laserX+gs.directionX < gs.laserSize {
		gs.directionX = -gs.directionX
	}

	if gs.laserY+gs.directionY > windowSize.h-gs.laserSize || gs.laserY+gs.directionY < gs.laserSize {
		gs.directionY = -gs.directionY
	}

	gs.laserX += gs.directionX
	gs.laserY += gs.directionY

	laserCtx.Call("clearRect", 0, 0, windowSize.w, windowSize.h)
	laserCtx.Call("beginPath")
	laserCtx.Call("arc", gs.laserX, gs.laserY, gs.laserSize, 0, math.Pi*2, false)
	laserCtx.Call("fill")
	laserCtx.Call("closePath")
}

func updatePlayer(event js.Value) {
	mouseX := event.Get("clientX").Float()
	mouseY := event.Get("clientY").Float()

	// basically threads/async/parallelism
	// TODO difference with Web Workers
	// TODO difference with Service Workers
	// https://gobyexample.com/goroutines
	go log("mouseEvent", "x", mouseX, "y", mouseY)

	// next gist
	if isLaserCaught(mouseX, mouseY, gs.laserX, gs.laserY) {
		go playSound()
	}
}

func setup() {
	window = js.Global()
	doc = window.Get("document")
	body = doc.Get("body")

	windowSize.h = window.Get("innerHeight").Float()
	windowSize.w = window.Get("innerWidth").Float()

	canvas = doc.Call("createElement", "canvas")
	canvas.Set("height", windowSize.h)
	canvas.Set("width", windowSize.w)
	body.Call("appendChild", canvas)

	laserCtx = canvas.Call("getContext", "2d")
	laserCtx.Set("fillStyle", "red")

	beep = window.Get("Audio").New("data:audio/mp3;base64,SUQzBAAAAAAAI1RTU0UAAAAPAAADTGF2ZjU2LjI1LjEwMQAAAAAAAAAAAAAA/+NAwAAAAAAAAAAAAFhpbmcAAAAPAAAAAwAAA3YAlpaWlpaWlpaWlpaWlpaWlpaWlpaWlpaWlpaWlpaWlpaW8PDw8PDw8PDw8PDw8PDw8PDw8PDw8PDw8PDw8PDw8PDw////////////////////////////////////////////AAAAAExhdmYAAAAAAAAAAAAAAAAAAAAAACQAAAAAAAAAAAN2UrY2LgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP/jYMQAEvgiwl9DAAAAO1ALSi19XgYG7wIAAAJOD5R0HygIAmD5+sEHLB94gBAEP8vKAgGP/BwMf+D4Pgh/DAPg+D5//y4f///8QBhMQBgEAfB8HwfAgIAgAHAGCFAj1fYUCZyIbThYFExkefOCo8Y7JxiQ0mGVaHKwwGCtGCUkY9OCugoFQwDKqmHQiUCxRAKOh4MjJFAnTkq6QqFGavRpYUCmMxpZnGXJa0xiJcTGZb1gJjwOJDJgoUJG5QQuDAsypiumkp5TUjrOobR2liwoGBf/X1nChmipnKVtSmMNQDGitG1fT/JhR+gYdCvy36lTrxCVV8Paaz1otLndT2fZuOMp3VpatmVR3LePP/8bSQpmhQZECqWsFeJxoepX9dbfHS13/////aysppUblm//8t7p2Ez7xKD/42DE4E5z9pr/nNkRw6bhdiCAZVVSktxunhxhH//4xF+bn4//6//3jEvylMM2K9XmWSn3ah1L2MqVIjmNlJtpQux1n3ajA0ZnFSu5EpX////uGatn///////1r/pYabq0mKT//TRyTEFNRTMuOTkuNaqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq/+MQxNIAAANIAcAAAKqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqg==")
}

func isLaserCaught(mouseX, mouseY, laserX, laserY float64) bool {
	return (math.Pow(mouseX-laserX, 2) + math.Pow(mouseY-laserY, 2)) < math.Pow(gs.laserSize+15, 2)
	// return laserCtx.Call("isPointInPath", mouseX, mouseY).Bool()
}

// no this isn't some magic; it's straight from HTML5
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLAudioElement#Basic_usage
func playSound() {
	beep.Call("play")
	window.Get("navigator").Call("vibrate", 300)
}

type gameState struct{ laserX, laserY, directionX, directionY, laserSize float64 }

// basically a rest+spread from javascript
// ...interface{} is more or less `any` from Typescript
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/rest_parameters#Description
func log(args ...interface{}) {
	window.Get("console").Call("log", args...)
}
