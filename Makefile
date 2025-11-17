MAIN_PACKAGE = ./cmd/app

.PHONY: run

run:
	@echo "Running $(MAIN_PACKAGE)..."
	go run $(MAIN_PACKAGE) $(ARGS)
