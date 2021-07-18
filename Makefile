
.PHONY: build
build:
	go build -o dist/calculator

.PHONY: deps
deps:
	go get -u github.com/onsi/ginkgo/ginkgo
	go get -u github.com/onsi/gomega

.PHONY: clean
clean:
	find . -name '*.coverprofile' -exec rm -f {} \;
	rm -rf dist/

.PHONY: test
test:
	ginkgo -r -cover --randomizeAllSpecs --randomizeSuites --race --trace **/*