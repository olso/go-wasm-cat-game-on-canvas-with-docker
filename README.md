# 🐈 go-wasm-cat-game-on-canvas-with-docker

Works with Chrome for Android! Thank you [@twifkak](https://github.com/twifkak)

Companion codebase to article [TODO add link]()

## Requirements
Docker or local Go 1.12-rc

## Setup Docker image if you need Chrome+Android
`docker build .` twifkak's golang fork

## Usage 🔧 💡
`make build_go` to build with Go compiler; produces __megabytes__ version 😿

`make build_twifkak` to build with golang fork which makes it work on Chrome+Android

`make serve` to run static http server

## Contact me
[https://olso.space](https://olso.space)

[https://twitter.com/olso_uznebolo](https://twitter.com/olso_uznebolo)

## Weird stuff

* WebAssembly is not enabled in iOS Simulator
* I had to use twifkak's Golang fork to instantiate in on phone, default Golang Wasm currently allocates too much memory
* For some reason, there is an `Audio` `play()` delay on Android
