run:
	bootdev run
sub:
	yes | bootdev run -s
go:
	go run .
test:
	go test ./...
