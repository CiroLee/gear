BIN=go

coverage:
	${BIN} test -v ./... -coverprofile=cover.out -covermode=atomic
	${BIN} tool cover -html=cover.out -o cover.html