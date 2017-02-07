FILES=main/*.go
TEST_FILES=tests/*.go

build:
	go run $(FILES)

test:
	go test $(TEST_FILES)
