BASE_FILE=main.go

init:
	go build && ./golb init

create:
	go build && ./golb create

parse:
	go build && ./golb parse