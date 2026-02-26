run: fmt check
	-rm $(HOME)/.config/todo.json
	go run . -add "Wash Fido"
	go run . -add "Wash car"
	go run . -add "Learn some Go"
	go run . -list
	go run . -toggle 0
	go run . -list
	go run . -delete 1
	go run . -list

fmt:
	go fmt .

check:
	go vet .
