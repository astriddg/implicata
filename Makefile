NAME = implicata

build:
	go build -o $(NAME)

test:
	go test ./... -cover
