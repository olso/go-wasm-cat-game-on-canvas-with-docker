# 🐈 go-wasm-cat-game-on-canvas-with-docker

Works with Chrome for Android! Thank you [@twifkak](https://github.com/twifkak)

Companion codebase to article

## Requirements
Docker, Go 1.12 (Release Candidate)

## Setup Docker image if you need Chrome+Android
`docker build .` twifkak's golang fork

## Usage 🔧 💡
`make build_go` to build with Go compiler; produces __megabytes__ version 😿

`make build_twifkak` to build with golang fork which makes it work on Chrome+Android

`make build`

`make serve` run nginx with proper WebAssembly __MIME__ type, Chrome/Firefox cries otherwise

## Contact me
[https://olso.space](https://olso.space)

[https://twitter.com/olso_uznebolo](https://twitter.com/olso_uznebolo)

## Weird stuff

* WebAssembly is not enabled in iOS Simulator
