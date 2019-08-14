build:
	go build -o takehome

run:
	./takehome
test:
	go test ./... -cover
test-html:
	go test ./... -coverprofile=cover.out && go tool cover -html=cover.out
vet:
	go vet -v ./...