async function fetchAndInstantiate(fileUrl, importObject) {
    try {
        const file = await fetch(fileUrl)
        const buffer = await file.arrayBuffer()
        const compiled = await WebAssembly.instantiate(buffer, importObject)

        return compiled.instance
    } catch (err) {
        console.error(err)
    }
}

async function run(fileUrl) {
    const go = new Go()
    const game = await fetchAndInstantiate(fileUrl, go.importObject)
    go.run(game)
}

setTimeout(async () => {
    run('game.wasm')
})
