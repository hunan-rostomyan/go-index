FILES=main/*.go
TEST_FILES=tests/*.go

all:
	go run $(FILES)

test:
	go test $(TEST_FILES)
