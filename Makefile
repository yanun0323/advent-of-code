.PHONY: run test
run:
	DAY=${DAY} go run main.go
test:
	go test --count=1 ./...