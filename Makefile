
.PHONY: run watch
run:
	@echo "Running..."
	go run . app:serve

watch:
	@echo "Watching..."
	air -d -c .air.toml