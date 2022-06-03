wire:
	cd cmd/blog && wire


run:
	go run cmd/blog/main.go cmd/blog/wire_gen.go