build:
	source "$(HOME)/dev/github.com/emscripten-core/emsdk/emsdk_env.sh" && \
	emcc hello.c -s WASM=1 -o dist/hello.html

http:
	open http://localhost:8000
	python -m http.server 8000
