.PHONY: init
init:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/mgechev/revive@v1.0.5


.PHONY: lint
lint:
	revive -config revive.toml -formatter unix pkg/...

.PHONY: fmt
fmt:
	goimports -e -w -d $(shell find ./pkg -type f -name '*.go' -print)
	gofmt -e -w -d $(shell find ./pkg -type f -name '*.go' -print)
