build: set GOARCH=amd64 set GOOS=${PLATFORM} go build -ldflags="-s -w" -o bin/${BINARY_NAME} ./...

run: ./${BINARY_NAME}

build_and_run: 
	build 
	run

deploy: serverless deploy

build_and_deploy: 
	build 
	serverless deploy

clean:
	go clean
	rm ${BINARY_NAME}