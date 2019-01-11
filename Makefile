rebuild:
	docker-compose up --build
build:
	docker-compose run go_wasm_build go build src/main.go -o src/main.wasm
run:
	open http://localhost:8080 && http-server ./server -p 8080
