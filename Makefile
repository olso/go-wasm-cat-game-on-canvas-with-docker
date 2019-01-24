build_twifkak:
	docker run --rm \
	-v `pwd`/src:/game \
	--env GOOS=js --env GOARCH=wasm --env GODEBUG=gcstoptheworld=1 --env GOGC=20 \
	olsansky/twifkak-go:latest \
	/bin/bash -c "go build -o /game/game.wasm /game/main.go; cp /usr/local/go/misc/wasm/wasm_exec.js /game/wasm_exec.js"

build_go:
	docker run --rm \
	-v `pwd`/src:/game \
	--env GOOS=js --env GOARCH=wasm \
	golang:1.12-rc \
	/bin/bash -c "go build -o /game/game.wasm /game/main.go; cp /usr/local/go/misc/wasm/wasm_exec.js /game/wasm_exec.js"

serve:
	docker run --rm -p 8080:8043 -v `pwd`/src:/srv/http pierrezemb/gostatic:latest
