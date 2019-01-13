build_tinygo:
	docker run \
	--rm \
	-v `pwd`:/game \
	tinygo/tinygo:latest \
    /bin/bash -c "tinygo build -o /game/dist/tinygo_game.wasm -target wasm /game/src/main.go; cp /go/src/github.com/aykevl/tinygo/targets/wasm_exec.js /game/dist/tinygo_wasm_exec.js"
build_go:
	docker run \
	--rm \
	-v `pwd`:/game \
	--env GOOS=js \
	--env GOARCH=wasm \
	golang:latest \
	/bin/bash -c "go build -o /game/dist/go_game.wasm /game/src/main.go; cp /usr/local/go/misc/wasm/wasm_exec.js /game/dist/go_wasm_exec.js"
build:
	make build_go
	make build_tinygo
