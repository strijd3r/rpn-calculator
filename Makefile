
.PHONY: build
build:
	go build -o dist/calculator

.PHONY: deps
deps:
	go mod download

.PHONY: clean
clean:
	find . -name '*.coverprofile' -exec rm -f {} \;
	rm -rf dist/

.PHONY: test
test:
	pre-commit run -a
	ginkgo -r -cover --randomizeAllSpecs --randomizeSuites --race --trace **/*