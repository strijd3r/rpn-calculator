
.PHONY: clean
clean:
	find . -name '*.coverprofile' -exec rm -f {} \;

.PHONY: test
test:
	ginkgo -r -cover -randomizeSuites **/*