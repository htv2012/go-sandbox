.PHONY: serve run

serve:
	go fmt ./...
	go vet ./...
	go run main.go

run:
	http GET localhost:8000/movies
	http PUT localhost:8000/movies/2 director:='{"firstname": "You", "lastname": "I"}'
	http GET localhost:8000/movies
