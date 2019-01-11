rebuild:
	docker-compose up -d --force-recreate --build go_cat_game_build
	docker-compose up -d --force-recreate --build tinygo_cat_game_build
go_build:
	docker-compose run go_cat_game_build go build -o build/go_game.wasm src/main.go
tinygo_build:
	docker-compose run tinygo_cat_game_build tinygo build -o build/tinygo_game.wasm -target wasm src/main.go
run:
	open http://localhost:8080 && http-server ./build -p 8080
