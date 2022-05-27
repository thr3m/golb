BASE_FILE=main.go

init:
	rm -rf ~/my_blog && go run ${BASE_FILE} init

create:
	go run ${BASE_FILE} create

parse:
	go build && ./golb parse