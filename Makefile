run: build
	./$< index.ssi

build: *.go go.mod
	go build -o $@ .