build_tinygo:
	docker run \
	--rm \
	-v `pwd`/src:/game \
	tinygo/tinygo:latest \
    /bin/bash -c "tinygo build -target wasm -o /game/tinygo_game.wasm /game/main.go; cp /go/src/github.com/aykevl/tinygo/targets/wasm_exec.js /game/tinygo_wasm_exec.js"

build_go:
	docker run \
	--rm \
	-v `pwd`/src:/game \
	--env GOOS=js \
	--env GOARCH=wasm \
	golang:latest \
	/bin/bash -c "go build -o /game/game.wasm /game/main.go; cp /usr/local/go/misc/wasm/wasm_exec.js /game/wasm_exec.js"

build:
	make build_go
	make build_tinygo
