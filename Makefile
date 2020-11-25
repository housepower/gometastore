PKG := github.com/akolb1/gometastore

GO        := go
GOBUILD   := $(GO) build $(BUILD_FLAG)

.PHONY: clean

clean:
	@rm -rf dist/*

preinstall: clean
	$(GO) mod tidy

build_hmsbench: preinstall
	$(GOBUILD) -o dist/hmsbench hmsbench/main.go

build_hmstool: preinstall
	$(GOBUILD) -o dist/hmstool hmstool/main.go

build_web: preinstall
	cd hmsweb ;\
	$(GOBUILD) -o ../dist/hmsweb .


build: build_hmsbench build_hmstool build_web
	@echo "building all"