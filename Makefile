build:
	go build -o godeck

test:
	# go test  ./... -cover -bench=
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	rm ./coverage.out

generate_completion: build
	./godeck completion zsh > /tmp/completion
