NAME = implicata

build:
	go build -o $(NAME)

run: build
	./$(NAME)

test:
	go test ./... -cover -race
