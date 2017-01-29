FILES=*.go
INDEX_JSON=index.json

all:
	go run $(FILES)

pretty:
	python prettify.py $(INDEX_JSON)

