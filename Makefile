build:
	 env GOOS=darwin go build -o bin/darwin/odin
	 env GOOS=linux  go build -o bin/linux/odin
install:
	go install
