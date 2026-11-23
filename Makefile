run:
	bootdev run
sub:
	yes | bootdev run -s
test:
	go test ./...
