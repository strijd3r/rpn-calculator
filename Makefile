
.PHONY: build
build:
	go build -o dist/calculator

.PHONY: clean
clean:
	find . -name '*.coverprofile' -exec rm -f {} \;
	rm -rf dist/

.PHONY: test
test:
	ginkgo -r -cover -randomizeSuites **/*