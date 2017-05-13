NAME = implicata

build:
	go build -race -o $(NAME)

run: build
	./$(NAME)

test:
	go test ./... -cover -race
