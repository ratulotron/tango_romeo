BINARY_NAME=tango-romeo

build:
	go build -o ${BINARY_NAME} main.go

run:
	./${BINARY_NAME} run

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}
